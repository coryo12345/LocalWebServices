import { ColumnDef, RowSelectionState } from "@tanstack/react-table";
import { FileMeta } from "localwebservices-sdk";
import { Search } from "lucide-react";
import { useMemo, useState } from "react";
import { addPath, progressState } from "../state";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { DataTable, sortableHeader } from "./ui/data-table";
import { Input } from "./ui/input";
import { connectionState } from "@/state";
import { toast } from "./ui/use-toast";
import { formatBytes } from "../lib/utils";

interface Props {
  files: FileMeta[] | null;
}

export function FileStorageTable({ files }: Props) {
  const [searchTxt, setSearchTxt] = useState("");
  const [selected, setSelected] = useState<RowSelectionState>({});

  const filteredFiles = useMemo(() => {
    return (
      files?.filter((f) => {
        return (
          f.name.toLowerCase().includes(searchTxt.toLowerCase()) ||
          f.lastModified.toLowerCase().includes(searchTxt.toLowerCase())
        );
      }) ?? []
    );
  }, [searchTxt, files]);

  async function downloadFiles() {
    const selectedFiles = files?.filter(
      (file, idx) => selected[idx] && file.realFile
    );
    if (!selectedFiles) return;

    progressState.value = 0;
    for (const idx in selectedFiles) {
      const file = selectedFiles[idx];
      const downloaded = await connectionState.client.downloadFile(
        file.relativePath,
        file.name
      );

      progressState.value = ((parseInt(idx) + 1) / selectedFiles.length) * 100;

      if (downloaded.hasError()) {
        toast({
          title: `Unable to download file: ${file.name}`,
          description: downloaded.getError(),
          variant: "destructive"
        });
      } else if (!downloaded.getValue()) {
        toast({
          title: `Unable to download file: ${file.name}`,
          description: "An unknown error has occured.",
          variant: "destructive"
        });
      }
    }
    progressState.value = -1;
  }

  return (
    <>
      <div className="flex justify-between mb-2">
        <div className="relative">
          <Input
            placeholder="Search"
            className="pl-8"
            value={searchTxt}
            onChange={(e) => setSearchTxt(e.target.value)}
          />
          <Search className="absolute top-1.5 left-1" />
        </div>
        <div>
          {/* TODO dropdown for download / delete */}
          <Button variant="outline" className="mr-2" onClick={downloadFiles}>
            Download
          </Button>
          {/* TODO upload */}
          <Button>Upload</Button>
        </div>
      </div>
      <DataTable
        columns={columns}
        data={filteredFiles}
        onSelectionChange={setSelected}
      />
    </>
  );
}

const columns: ColumnDef<FileMeta>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="block"
      />
    ),
    cell: ({ row }) => {
      if (row.original.name === "..") {
        return <span></span>;
      }
      return (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
          className="block"
        />
      );
    },
    enableSorting: false,
    enableHiding: false
  },
  {
    accessorKey: "name",
    header: sortableHeader("Filename"),
    cell: ({ row }) => {
      const isDir = row.original.isDir;
      const name: string = row.original.name;
      const displayName = name === ".." ? name : name + "/";
      if (isDir) {
        return (
          <span
            className="cursor-pointer hover:underline text-primary pr-4"
            onClick={() => addPath(name)}
          >
            {displayName}
          </span>
        );
      } else {
        return <span>{name}</span>;
      }
    }
  },
  {
    accessorKey: "size",
    header: sortableHeader("Size"),
    cell: ({ row }) => {
      if (!row.original.realFile) return <span></span>;
      const size = formatBytes(row.original.size);
      return <span>{size}</span>;
    }
  },
  {
    accessorKey: "lastModified",
    header: sortableHeader("Last Modified"),
    cell: ({ row }) => {
      if (!row.original.realFile) return <span></span>;
      return (
        <span>{new Date(row.original.lastModified).toLocaleString()}</span>
      );
    }
  }
];

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from "@/components/ui/table";
import {
  ColumnDef,
  HeaderContext,
  RowSelectionState,
  SortingState,
  Updater,
  flexRender,
  getCoreRowModel,
  getSortedRowModel,
  useReactTable
} from "@tanstack/react-table";
import { ChevronDown, ChevronUp } from "lucide-react";
import { useEffect, useState } from "react";
import { Button } from "./button";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data?: TData[] | null;
  noDataMessage?: string;
  onSelectionChange?: (state: RowSelectionState) => void;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  ...props
}: DataTableProps<TData, TValue>) {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [rowSelection, setSelection] = useState<RowSelectionState>({});

  function selectionUpdate(state: Updater<RowSelectionState>) {
    if (
      props.onSelectionChange &&
      typeof props.onSelectionChange === "function"
    ) {
      if (typeof state === "function") {
        props.onSelectionChange(state(rowSelection));
      } else {
        props.onSelectionChange(state);
      }
    }
    setSelection(state);
  }

  useEffect(() => {
    selectionUpdate({});
  }, [data]);

  const table = useReactTable({
    data: data ?? [],
    columns,
    getCoreRowModel: getCoreRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onRowSelectionChange: selectionUpdate,
    state: {
      sorting,
      rowSelection
    }
  });

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                {props.noDataMessage ?? "No results."}
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}

// eslint-disable-next-line react-refresh/only-export-components
export const sortableHeader = (name: string) =>
  function <TData>({ column }: HeaderContext<TData, unknown>) {
    return (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        className="w-full text-start justify-start px-0 hover:bg-inherit"
      >
        {name}
        {column.getIsSorted() === "asc" && (
          <ChevronUp className="ml-2 h-4 w-4" />
        )}
        {column.getIsSorted() === "desc" && (
          <ChevronDown className="ml-2 h-4 w-4" />
        )}
      </Button>
    );
  };

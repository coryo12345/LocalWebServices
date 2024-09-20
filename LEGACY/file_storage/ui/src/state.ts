import { toast } from "@/components/ui/use-toast";
import { FileMeta, FileStorage } from "localwebservices-sdk";
import path from "path";
import { proxy } from "valtio";
import { watch } from "valtio/utils";

export const connectionState = proxy({
  url: "http://localhost:8081/",
  get client() {
    return FileStorage(connectionState.url);
  }
});

export const pathState = proxy({
  pathItems: ["."],
  get pathDisplay(): string {
    const items = pathState.pathItems;
    return path.join(...items, "/");
  }
});

export function addPath(name: string) {
  if (name === ".." || name === "../") {
    pathState.pathItems.pop();
  } else {
    pathState.pathItems.push(name);
  }
}

interface FileState {
  files: FileMeta[] | null;
}

export const fileState = proxy<FileState>({
  files: []
});

watch((get) => {
  const client = get(connectionState).client;
  const currentPath = get(pathState).pathDisplay;

  const abortController = new AbortController();

  (async function () {
    const filesResult = await client.listFiles(currentPath, {
      signal: abortController.signal
    });
    if (filesResult.hasError()) {
      toast({
        title: "Something went wrong...",
        description: "Unable to fetch files",
        variant: "destructive"
      });
      fileState.files = null;
    } else {
      const _files = filesResult.getValue();
      _files.sort(filemetaComparator);
      fileState.files = _files;
    }
  })();

  return () => {
    abortController.abort();
  };
});

function filemetaComparator(a: FileMeta, b: FileMeta): number {
  if (a.isDir && !b.isDir) return -1;
  else if (!a.isDir && b.isDir) return 1;
  else return a.name.localeCompare(b.name);
}

export const progressState = proxy({
  value: -1
});

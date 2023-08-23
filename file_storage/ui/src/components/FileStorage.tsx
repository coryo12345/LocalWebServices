import { useSnapshot } from "valtio";
import { FileStorageTable } from "./FileStorageTable";
import { fileState, pathState } from "../state";

export function FileStorage() {
  const pathSnap = useSnapshot(pathState);
  useSnapshot(fileState);
  return (
    <>
      <hr className="mb-2" />
      <div className="ml-1 mb-2 text-sm">
        Showing files for:{" "}
        <code className="px-2 py-1 rounded bg-secondary">
          {pathSnap.pathDisplay}
        </code>
      </div>
      <FileStorageTable files={fileState.files} />
    </>
  );
}

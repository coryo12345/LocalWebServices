import { useContext } from "preact/hooks";
import { DataTable } from "./DataTable";
import { AppState } from "../state";
import { useSignal } from "@preact/signals";
import { ListAction } from "../models";
import { FilterBar } from "./FilterBar";
import type { Property } from "localwebservices-sdk";

export function PropertyList() {
  const state = useContext(AppState);

  const selectedProperties = useSignal<Property[]>([]);

  function handleAction(action: ListAction) {
    // TODO: handle actions using SDK
    console.log(action);
  }

  return (
    <>
      <FilterBar onAction={handleAction} />
      <div class="mx-2">
        <DataTable
          items={state.filteredProperties.value}
          onSelect={(items) => (selectedProperties.value = items)}
        />
      </div>
    </>
  );
}

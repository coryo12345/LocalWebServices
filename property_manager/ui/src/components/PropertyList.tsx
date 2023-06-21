import { useContext } from "preact/hooks";
import { DataTable } from "./DataTable";
import { AppState } from "../state";
import { useSignal } from "@preact/signals";
import { ListAction, Property } from "../models";
import { FilterBar } from "./FilterBar";

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
      <DataTable
        items={state.filteredProperties.value}
        onSelect={(items) => (selectedProperties.value = items)}
      />
    </>
  );
}

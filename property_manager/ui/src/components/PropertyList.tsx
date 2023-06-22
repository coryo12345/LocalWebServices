import { useSignal } from "@preact/signals";
import type { Property } from "localwebservices-sdk";
import { useContext } from "preact/hooks";
import { AppState } from "../state";
import { FilterBar } from "./FilterBar";
import { DataTable, ListAction } from "./base/DataTable";
import { ADD_PROPERTY_EVENT, AddPropertyDialog } from "./AddPropertyDialog";
import { useEventBus } from "../utils/bus";

export function PropertyList() {
  const state = useContext(AppState);
  const { emit } = useEventBus();

  const selectedProperties = useSignal<Property[]>([]);

  function handleAction(action: ListAction) {
    switch (action) {
      case "add":
        emit(ADD_PROPERTY_EVENT);
        break;
      case "delete":
        // use selectedProperties to delete items
        break;
      case "update":
        // use selectedProperties to update values
        break;
      default:
        break;
    }
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
      <AddPropertyDialog />
    </>
  );
}

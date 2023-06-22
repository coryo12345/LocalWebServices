import { useSignal } from "@preact/signals";
import type { Property } from "localwebservices-sdk";
import { useContext } from "preact/hooks";
import { AppState } from "../state";
import { events, useEventBus } from "../utils/bus";
import { ADD_PROPERTY_EVENT, AddPropertyDialog } from "./AddPropertyDialog";
import { FilterBar } from "./FilterBar";
import { DataTable, ListAction } from "./base/DataTable";

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
        deleteProperties();
        break;
      case "update":
        // TODO use selectedProperties to update values
        break;
      default:
        break;
    }
  }

  async function deleteProperties() {
    const keys = selectedProperties.value.map((prop) => prop.key);
    const result = await state.deleteProperties(keys);
    if (result) {
      emit(events.SNACKBAR_SUCCESS, "Successfully deleted properties");
    } else {
      emit(events.SNACKBAR_ERROR, "Failed to delete 1 or more properties");
    }
    selectedProperties.value = [];
    state.fetchProperties();
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

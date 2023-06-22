import { useSignal } from "@preact/signals";
import type { Property } from "localwebservices-sdk";
import { useContext } from "preact/hooks";
import { AppState } from "../state";
import { events, useEventBus } from "../utils/bus";
import { ADD_PROPERTY_EVENT, AddPropertyDialog } from "./AddPropertyDialog";
import {
  UPDATE_PROPERTY_EVENT,
  UpdatePropertyDialog,
} from "./UpdatePropertyDialog";
import { FilterBar } from "./FilterBar";
import { DataTable, ListAction } from "./base/DataTable";

export function PropertyList() {
  const state = useContext(AppState);
  const { emit } = useEventBus();

  const selectedProperties = useSignal<Property[]>([]);

  function handleAction(action: ListAction) {
    const keys = selectedProperties.value.map((prop) => prop.key);
    switch (action) {
      case "add":
        emit(ADD_PROPERTY_EVENT);
        break;
      case "delete":
        deleteProperties(keys);
        break;
      case "update":
        emit(UPDATE_PROPERTY_EVENT, keys);
        break;
      default:
        break;
    }
  }

  async function deleteProperties(keys: string[]) {
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
      <UpdatePropertyDialog />
    </>
  );
}

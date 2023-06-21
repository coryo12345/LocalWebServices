import { computed, signal, useSignalEffect } from "@preact/signals";
import { Property } from "./models";
import { createContext } from "preact";
import { PropertyManager } from "localwebservices-sdk";

export function createState() {
  const url = signal("http://localhost:8081");
  const client = signal({} as ReturnType<typeof PropertyManager>);
  useSignalEffect(() => {
    client.value = PropertyManager(url.value);
  });

  const searchText = signal("");
  const searchKey = signal(true);
  const searchValue = signal(false);
  const properties = signal<Property[] | null>([]);

  const filteredProperties = computed(() => {
    const s = searchText.value.toLowerCase();
    return (
      properties.value?.filter(
        (p) =>
          (searchKey.value && p.key.toLowerCase().includes(s)) ||
          (searchValue.value && p.value.toLowerCase().includes(s))
      ) ?? []
    );
  });

  async function fetchProperties() {
    properties.value = await client.value.getAllProperties();
  }

  return {
    url,
    searchText,
    searchKey,
    searchValue,
    properties,
    filteredProperties,
    fetchProperties,
  };
}

export const AppState = createContext({} as ReturnType<typeof createState>);

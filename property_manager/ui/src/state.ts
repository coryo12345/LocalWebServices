import { computed, signal, useSignalEffect } from "@preact/signals";
import { createContext } from "preact";
import { Property, PropertyManager } from "localwebservices-sdk";

export function createState() {
  const url = signal("http://localhost:8081");

  // ideally this client stuff would be moved into a separate service-esque file
  // and the url / connection handling would be much more graceful.
  // but the scope of this is small enough to just keep it here for now
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

  async function addProperty(
    key: string,
    value: string
  ): Promise<string | null> {
    return await client.value.setProperty(key, value);
  }

  return {
    url,
    searchText,
    searchKey,
    searchValue,
    properties,
    filteredProperties,
    fetchProperties,
    addProperty,
  };
}

export const AppState = createContext({} as ReturnType<typeof createState>);

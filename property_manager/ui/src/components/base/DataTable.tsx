import { useComputed, useSignal } from "@preact/signals";
import { Property } from "localwebservices-sdk";
import { useEffect } from "preact/hooks";

export type ListAction = "delete" | "add" | "update";

interface Item extends Property {
  selected?: boolean;
}

interface Props {
  items: Item[];
  onSelect?: (items: Item[]) => void;
}

// TODO column sorting
export function DataTable(props: Props) {
  const _items = useSignal<Item[]>([]);

  useEffect(() => {
    _items.value = structuredClone(props.items) ?? [];
  }, [props.items]);

  function selectHandler(index: number) {
    _items.value[index].selected = !_items.value[index].selected;
    typeof props.onSelect === "function" &&
      props.onSelect(_items.value.filter((i) => i.selected));
  }

  const sortBy = useSignal("key");
  const sortAsc = useSignal(true);
  function setSort(key: string) {
    if (key === sortBy.value && sortAsc.value) {
      sortAsc.value = false;
    } else if (key === sortBy.value) {
      sortBy.value = "";
    } else {
      sortBy.value = key;
      sortAsc.value = true;
    }
  }
  function getSortIcon(key: string) {
    if (key === sortBy.value && sortAsc.value) {
      return <span>&#9650;</span>;
    } else if (key === sortBy.value) {
      return <span>&#9660;</span>;
    } else {
      return <span></span>;
    }
  }

  const sortedItems = useComputed(() => {
    return [..._items.value].sort((a, b) => {
      if (sortBy.value === "") return 0;
      const aVal = (a as Record<string, any>)[sortBy.value] as string;
      const bVal = (b as Record<string, any>)[sortBy.value] as string;
      return sortAsc.value
        ? aVal.localeCompare(bVal)
        : bVal.localeCompare(aVal);
    });
  });

  return (
    <table class="table-auto w-full">
      <thead className="bg-slate-200 border border-black">
        <tr>
          <th></th>
          <th className="py-1 px-2" onClick={() => setSort("key")}>
            Key {getSortIcon("key")}
          </th>
          <th className="py-1 px-2" onClick={() => setSort("value")}>
            Value {getSortIcon("value")}
          </th>
        </tr>
      </thead>
      <tbody>
        {sortedItems.value.map((item, index) => (
          <tr className="border border-black hover:bg-slate-100">
            <td>
              <input
                checked={!!item.selected}
                type="checkbox"
                className="text-center mx-2 w-4 h-4"
                onChange={() => selectHandler(index)}
              />
            </td>
            <td className="text-center py-1 px-2">{item.key}</td>
            <td className="text-center py-1 px-2">{item.value}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

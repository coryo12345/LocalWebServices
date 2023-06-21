import { useSignal } from "@preact/signals";
import { Property } from "localwebservices-sdk";
import { useEffect } from "preact/hooks";

interface Item extends Property {
  selected?: boolean;
}

interface Props {
  items: Item[];
  onSelect?: (items: Item[]) => void;
}

export function DataTable(props: Props) {
  const _items = useSignal<Item[]>([]);

  useEffect(() => {
    _items.value = structuredClone(props.items) ?? [];
  }, [props.items]);

  function selectHandler(index: number) {
    _items.value[index].selected = !_items.value[index].selected;
  }

  return (
    <table class="table-auto w-full">
      <thead className="bg-slate-200 border border-black">
        <tr>
          <th></th>
          <th className="py-1 px-2">Key</th>
          <th className="py-1 px-2">Value</th>
        </tr>
      </thead>
      <tbody>
        {_items.value.map((item, index) => (
          <tr className="border border-black hover:bg-slate-100">
            <td>
              <input
                checked={item.selected}
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

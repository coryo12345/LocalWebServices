import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { connectionState } from "../state";

type Props = React.HTMLAttributes<HTMLDivElement>;

export function ConenctionControl({ className, ...props }: Props) {
  const [value, setValue] = useState(connectionState.url);
  return (
    <div className={cn("flex", className)} {...props}>
      <Input
        value={value}
        placeholder="http://localhost:8081"
        className="w-[20em] mr-2"
        onChange={(e) => setValue(e.target.value)}
      />
      <Button
        onClick={() => {
          connectionState.url = value;
        }}
      >
        Connect
      </Button>
    </div>
  );
}

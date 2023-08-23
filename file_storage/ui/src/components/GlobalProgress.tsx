import { useSnapshot } from "valtio";
import { Progress } from "./ui/progress";
import { progressState } from "../state";

export function GlobalProgress() {
  const progressSnap = useSnapshot(progressState);

  const showProgress = progressSnap.value >= 0 && progressSnap.value <= 100;
  return (
    <div className="fixed top-0 left-0 w-full">
      {showProgress && <Progress value={progressSnap.value} />}
    </div>
  );
}

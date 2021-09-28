import { PartToRename } from "@/pb/baseService_pb";

export default function convertNamesList(
  basicNamesList: PartToRename[]
): PartToRename[] {
  const tasks: PartToRename[] = [];
  let finished = false;
  while (!finished) {
    finished = true;
    basicNamesList.forEach((part) => {
      if (!finished) {
        return false;
      }
      // console.log(`iterate basicNamesList: old=${oldPath} | new=${newPath} | tasks=`, tasks)
      if (part.getOldname() !== part.getNewname()) {
        tasks.push(getReplacement(part.getOldname(), part.getNewname()));
        replacePaths(basicNamesList, tasks);
        finished = false;
        return false;
      }
    });
  }
  // console.log('final tasks=', tasks)
  return tasks;
}

function getReplacement(oldPath: string, newPath: string): PartToRename {
  const oldPathSegments = oldPath.split("/");
  const newPathSegments = newPath.split("/");
  for (let i = 0; i < oldPathSegments.length; i++) {
    if (oldPathSegments[i] !== newPathSegments[i]) {
      const part = new PartToRename();
      part.setOldname(oldPathSegments.filter((v, j) => j <= i).join("/"));
      part.setNewname(newPathSegments[i]);
      return part;
    }
  }

  throw new Error("couldn't make replacement: " + oldPath + " => " + newPath);
}

function replacePaths(
  basicNamesList: PartToRename[],
  tasks: PartToRename[]
): void {
  // console.log('replacePaths', basicNamesList, tasks)
  tasks.forEach((part) => {
    const oldname = part.getOldname();
    let newname = part.getNewname();
    // console.log(`check ${oldPath} to replace with ${newPath}`)
    const oldPathSegments = oldname.split("/");
    oldPathSegments.pop();
    newname =
      (oldPathSegments.length === 0 ? "" : oldPathSegments.join("/") + "/") +
      newname;
    basicNamesList.forEach((v, i) => {
      // console.log(`check ${basicNamesList[i][0]} for ${oldPath} to replace with ${newPath}`)
      if (basicNamesList[i].getOldname().indexOf(oldname) === 0) {
        basicNamesList[i].setOldname(
          newname + basicNamesList[i].getOldname().substr(oldname.length)
        );
        // console.log('replaced to ' + basicNamesList[i][0])
      } else {
        // console.log('not replaced')
      }
    });
  });
}

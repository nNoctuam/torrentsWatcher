
export default function convertNamesList (basicNamesList) {
  const tasks = []
  let finished = false
  while (!finished) {
    finished = true
    basicNamesList.forEach(([oldPath, newPath]) => {
      if (!finished) {
        return false
      }
      // console.log(`iterate basicNamesList: old=${oldPath} | new=${newPath} | tasks=`, tasks)
      if (oldPath !== newPath) {
        tasks.push(getReplacement(oldPath, newPath))
        replacePaths(basicNamesList, tasks)
        finished = false
        return false
      }
    })
  }
  // console.log('final tasks=', tasks)
  return tasks
}

function getReplacement (oldPath, newPath) {
  const oldPathSegments = oldPath.split('/')
  const newPathSegments = newPath.split('/')
  for (let i = 0; i < oldPathSegments.length; i++) {
    if (oldPathSegments[i] !== newPathSegments[i]) {
      return [
        oldPathSegments.filter((v, j) => j <= i).join('/'),
        newPathSegments[i]
      ]
    }
  }

  throw new Error('couldn\'t make replacement: ' + oldPath + ' => ' + newPath)
}

function replacePaths (basicNamesList, tasks) {
  // console.log('replacePaths', basicNamesList, tasks)
  tasks.forEach(([oldPath, newPath]) => {
    // console.log(`check ${oldPath} to replace with ${newPath}`)
    const oldPathSegments = oldPath.split('/')
    oldPathSegments.pop()
    newPath = (oldPathSegments.length === 0 ? '' : oldPathSegments.join('/') + '/') + newPath
    basicNamesList.forEach((v, i) => {
      // console.log(`check ${basicNamesList[i][0]} for ${oldPath} to replace with ${newPath}`)
      if (basicNamesList[i][0].indexOf(oldPath) === 0) {
        basicNamesList[i][0] = newPath + basicNamesList[i][0].substr(oldPath.length)
        // console.log('replaced to ' + basicNamesList[i][0])
      } else {
        // console.log('not replaced')
      }
    })
  })

  return basicNamesList
}

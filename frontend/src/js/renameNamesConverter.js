
export default function convertNamesList (basicNamesList) {
  const replaceMap = {}
  const replaceTasks = []
  Object.keys(basicNamesList).forEach(k => {
    Object.keys(replaceMap).forEach(key => {
      if (k.indexOf(key) === 0) {
        k = replaceMap[key] + k.substr(0, key.length)
        // console.log('replace oldPath at start: ' + key + ' => ' + k.split('/'))
      }
    })
    let oldPath = k.split('/')
    const newPath = basicNamesList[k].split('/')

    oldPath.forEach((segment, i) => {
      if (i < oldPath.length && segment !== newPath[i]) {
        const key = oldPath.filter((v, j) => j <= i).join('/')
        if (replaceMap[key] && replaceMap[key] !== newPath[i]) {
          throw new Error('path renamed differently: ' +
            'old=' + key +
            'renamed=' + replaceMap[key] +
            'renamed=' + newPath[i]
          )
        }
        replaceMap[key] = newPath.filter((v, j) => j <= i).join('/')
        replaceTasks.push([key, newPath[i]])

        Object.keys(replaceMap).forEach(key => {
          if (k.indexOf(key) === 0) {
            // console.log('replace oldPath in iteration: ' + k + ' => ' + (replaceMap[key] + k.substr(key.length)) + ' || key=' + key)
            k = replaceMap[key] + k.substr(key.length)
            oldPath = k.split('/')
          }
        })
      }
    })
  })

  return replaceTasks
}

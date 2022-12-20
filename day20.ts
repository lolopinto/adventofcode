import * as fs from "fs"

function day20(key: number, times: number) {

  const lines = fs.readFileSync("day20input").toString().split("\n");
  const vals:number[] = [];
  const indices: number[]=[];
  let i = 0;
  for (const line of lines) {
    vals.push(parseInt(line)*key);
    indices.push(i);
    i++;
  }

  for (let k=0; k < times;k++) {
    for (let i =0;i<indices.length;i++) {
      const curr_idx =indices.indexOf(i);
      indices.splice(curr_idx,1);
      let new_idx = (curr_idx + vals[i]) % indices.length;
      // console.log('new idx',i, curr_idx, vals[i],new_idx)
      indices.splice(new_idx,0, i)
      // console.log(indices)
    }
  }
  // console.log(indices)

  let zeroIndex = indices.indexOf(vals.indexOf(0))

  // console.log(zeroIndex)
  let sum = 0;
  for (const v of [1000,2000,3000]) {
    // console.log(vals[indices[(zeroIndex+v) % vals.length]])
    sum+=  vals[indices[(zeroIndex+v) % vals.length]]
  }
  return sum
}

// part 1
console.log(day20(1,1))
// part 2
console.log(day20(811589153,10))
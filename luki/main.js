const fs = require('fs')
 
try {
  var data = fs.readFileSync('./input.txt', 'utf8');
  inputData = data   
} catch(e) {
  console.log('Error:', e.stack);
}
 
inputData = inputData.split("\n")
 
 
let busIds = inputData[1].split(",")
console.log(busIds.toString());
 
let inLoop = true;
let i = 1;
let departTimeModifier = 1
let departTime = Number(busIds[0]);
let departIncrementor = departTime;
 
while (inLoop){
     
  if (i == busIds.length){
    inLoop = false;
    continue
  }
  if (busIds[i] === "x"){
        departTimeModifier+=1;
        i+=1
        continue
  }
  departTime += departIncrementor;
 
  let newDepartTime = departTime + departTimeModifier;
 

  if (newDepartTime % Number(busIds[i]) == 0){
    //success! found one
    departTime = newDepartTime - departTimeModifier
    departIncrementor = departIncrementor * Number(busIds[i])
    i+=1
    departTimeModifier+=1;
 
  }

}
console.log(departTime)
 
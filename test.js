




function A() {
  console.log("called A");
  return false;
}

function B() {
  console.log("called B");
  return true;
}

console.log(B() || console.log("A"));



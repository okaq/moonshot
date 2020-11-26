// kinetic curves
// created by aq <aq@okaq.com>
// 2020-11-25

// globals
const w0 = 1024;
const h0 = 1024;

const w1 = 32;
const h1 = 32;

const nx = w0 / w1;
const ny = h0 / h1;
const nt = nx * ny;
console.log(nx,ny,nt);

function setup() {
    // let w0 = 1024;
    // let h0 = 1024;
    createCanvas(w0, h0);
    noLoop();
  }
  
  function draw() {
    background(128);
    grid();
  }
  
  function grid() {
    strokeWeight(2);
    stroke(0);
    // x-y loop
  }
  
  
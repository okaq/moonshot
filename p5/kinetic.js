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

let line0 = [];
let lineN = 128;

function setup() {
    // let w0 = 1024;
    // let h0 = 1024;
    createCanvas(w0, h0);
    noLoop();
  }
  
  function draw() {
    background(128);
    // grid();
    zing();
  }
  
  function grid() {
    strokeWeight(2);
    stroke(0);
    // x-y loop
    for (let x0 = 0; x0 < nx; x0++) {
        let x1 = x0 * w1;
        line(0,x1,w0,x1);
      }
      for (let y0 = 0; y0 < ny; y0++) {
        let y1 = y0 * h1;
        line(y1,0,y1,h0);
      }
  }
  
  function zing() {
    gen();
  }

  function gen() {
    for (let i = 0; i < lineN; i++) {
      let line1 = [random(w0),random(h0)];
      line0.push(line1);
    }
  }
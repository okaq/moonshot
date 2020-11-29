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

let tick = 0;
let max2 = 16;
let id = null;

function setup() {
    // let w0 = 1024;
    // let h0 = 1024;
    createCanvas(w0, h0);
    noLoop();
  }
  
  function draw() {
    background(8);
    // grid();
    zing();
    // loop
    // use setInterval
    // loop();
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
    ren();
    loop2();
  }

  function gen() {
    for (let i = 0; i < lineN; i++) {
      let line1 = [];
      line1[0] = random(w0);
      line1[1] = random(h0);
      line1[2] = random(w0);
      line1[3] = random(h0);
      line0.push(line1);
    }
  }

  function ren() {
    stroke(20,200,10);
    for (let i = 0; i < lineN; i++) {
      let p0 = line0[i];
      let x0 = p0[0];
      let y0 = p0[1];
      let x1 = p0[2];
      let y1 = p0[3];
      line(x0,y0,x1,y1);
    }
  }

function loop2() {
  id = setInterval(frame, 1000);
}

function frame() {
  console.log("tick: " + tick);
  if (tick >= max2) {
    clearInterval(id);
    console.log("anim done.");
    return;
  }
  tick = tick + 1;
}


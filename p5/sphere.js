// sphere space division algo
// created by aq <aq@okaq.com>
// 2020-11-02

function setup() {
  let w0 = 3840;
  let h0 = 2160;
  createCanvas(w0, h0);
  noLoop()
}

function draw() {
  background(128,128,128);
  // grid
  strokeWeight(2);
  stroke(0)
  // rand pt
  let x0 = random(3840);
  let y0 = random(2160);
  let x1 = 3840;
  let y1 = 2160;
  line(x0,y0,x1,y1);
}

function randp(w0, h0) {
  let x0 = random(w0);
  let y0 = random(h0);
  return [x0,y0];
}



//helpers to help with understanding the morphling test replay
//
//takes a tick and gives a time
var toTime = function(x) {
  var sec = (x - 4220) / 30;
  var min = sec / 60 - (sec % 60) / 60;
  return "time: " + (sec < 0 ? "-" : "") + Math.abs(min) + ":" + Math.round(Math.abs(sec % 60))
}

var toTimeEmitter = function(tick,pretime,gametime) {
  var nowtime = gametime - pretime - 90;
  var tockoffset = tick - (nowtime * 30)
  return tockoffset
}

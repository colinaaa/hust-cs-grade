type Huster = {
  XH: string; //学号
  XM: string; //姓名
  DWMC: string; // 学院（单位名称）
};
const Husters = require("../data/husters.json") as Huster[];

const cs18 = Husters.filter(
  item => item.DWMC.indexOf("计算机") != -1 && item.XH.indexOf("U2018") != -1
).map(old => {
  return { XH: old.XH, XM: old.XM, DWMC: old.DWMC };
});

import { writeFile } from "fs";

writeFile("./data/cs18.json", JSON.stringify(cs18, null, "  "), err =>
  console.log(err ? err : "done!")
);

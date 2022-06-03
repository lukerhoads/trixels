// ZPP - Zoom Pan Pinch
// Functionality:
// - Zoom / shrink with button clicks

import { Pixel } from "types/pixel"

let canvas = document.getElementById("canvas") as HTMLCanvasElement
let ctx = canvas.getContext("2d") // remove





// export const draw = (pixels: Pixel[]) => {
//     if (!ctx) {
//         return
//     }

//     canvas.width = window.innerWidth
//     canvas.height = window.innerHeight

//     ctx.translate(window.innerWidth / 2, window.innerHeight / 2)
//     ctx.scale(scale, scale)

//     ctx.fillStyle = "#000"
//     ctx.fillRect(-50, -50, 100, 100)
// }
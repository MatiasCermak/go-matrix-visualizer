import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

export default async function calculateMatrix(rows, columns) {
    try {
        const res = await fetch(`http://localhost:8080/matrix/spiralFibonacci?rows=${rows}&cols=${columns}`);
        const data = await res.json();
        return data;
    } catch (err) {
        alert(err);
    }
}

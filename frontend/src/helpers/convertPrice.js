import axios from "axios"
import chains from '../chains/chains.json'

export const getPrice = async () => {
    let ids = ''
    Object.keys(chains).map(key => ids += chains[key].coingecko + ',')
    const res = await axios.get(`https://api.coingecko.com/api/v3/simple/price?ids=${ids}&vs_currencies=usd`)
    if(res.status === 200){
        return res.data
    }
}

export const convertPrice = async (amount, coingecko) => {
    const prices = await getPrice()
    const price = prices[coingecko]['usd'] * amount
    return price 
}


export const getTotal = async (balances) => {
    let sum = 0
    for ( let i of balances){
        const price = await convertPrice(parseFloat(i.amount), i.coingecko)
        sum += price
    }
    return sum.toPrecision(5)
}
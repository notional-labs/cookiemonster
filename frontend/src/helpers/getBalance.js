import axios from "axios";
import { getKeplr } from "./getKeplr";
import chains from '../chains/chains.json'
import ibc from '../chains/mapIBC.json'

const osmosisLogo = 'https://dl.airtable.com/.attachments/4ef30ec4008bc86cc3c0f74a6bb84050/0eeb4d64/aQ5W3zaT_400x400.jpg'
const ionLogo = 'https://app.osmosis.zone/public/assets/tokens/ion.png'

const mapIbc = (balance) => {
    let returnBalance = {
        denom: '',
        name: '',
        amount: '',
        logo: ''
    }
    if(balance.denom.substring(0,4) === "ibc/"){
        const key = ibc[`${balance.denom}`]
        if(chains[`${key}`]) {
            returnBalance.denom = chains[`${key}`]['denom']
            const chain_name = chains[`${key}`]['chain_name'] 
            returnBalance.name = key + ' - ' + chain_name.toUpperCase() 
            returnBalance.amount = (parseInt(balance.amount)/1000000).toString()
            returnBalance.logo = chains[`${key}`]['logo']
        }
    }
    else {
        returnBalance.denom = balance.denom
        const chain_name = balance.denom === 'uosmo' ? 'osmo' : 'ion'
        returnBalance.name = chain_name.toUpperCase()
        returnBalance.amount = (parseInt(balance.amount)/1000000).toString()
        returnBalance.logo = balance.denom === 'uosmo' ? osmosisLogo : ionLogo
    }
    return returnBalance
} 

export const getBalance = async (apiUrl = "https://lcd-osmosis.keplr.app", address) => {
    const URL = `${apiUrl}/bank/balances/${address}`
    let balances = []
    const res= await axios.get(URL)
    if(res.data && res.data.result){
        res.data.result.map(b => {
            const newBalance = mapIbc(b)
            balances.push(newBalance)
        })   
    }
    return balances
}

export const getOsmo = async (address) => {
    const balances = await getBalance(undefined, address)
    if(balances.length === 0){
        return 0
    }
    const filterBalance = balances.filter(x => x.denom === 'uosmo')
    return filterBalance[0].amount
}

export const getKeplrBalances = () => {
    let balances = []
    Object.keys(chains).map( async (chain) => {
        const { accounts } = await getKeplr(chains[chain].chain_id);
        const { data } = getBalance(chains[chain].apiUrl, accounts[0].address)
        if (data && data.result && data.result.length > 0){
            balances.push({
                balance: data.result,
                logo: chains[chain].logo,
            })
        }
    })
    return balances
}
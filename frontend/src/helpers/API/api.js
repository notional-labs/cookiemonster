import axios from "axios";

const root_Url = 'http://localhost:8000'

export const checkAccount = async (address) => {
    const res = await axios.post(`${root_Url}/check-account`, { address })
    console.log(res)
    return res
}

export const deposit = async (address, txHash) => {
    const res = await axios.post(`${root_Url}/deposit`, { address, 'tx-hash': txHash })
    return res
}

export const autoInvest = async (address) => {
    const res = await axios.post(`${root_Url}/auto-invest`, { address })
    return res
}

export const pullReward = async (address) => {
    const res = await axios.post(`${root_Url}/pull-reward`, { address })
    return res
}

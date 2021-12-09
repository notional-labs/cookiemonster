import axios from "axios"

export const checkReward = async (cookieMonster) => {
    const res = await axios.get(`https://lcd-osmosis.keplr.app/cosmos/distribution/v1beta1/delegators/${cookieMonster}/rewards`)
    return res
}
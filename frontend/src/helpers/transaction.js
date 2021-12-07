import { coins } from "@cosmjs/launchpad";

export const transaction = async (cosmJS, amount = 1000000) => {
    // define memo (not required)
    const memo = "Deposit";
    // sign and broadcast Tx
    // THIS IS A TEST ADDRESS 
    const recipient = "osmo1cptdzpwjc5zh6nm00dvetlg24rv9j3tjh7wnnz";
    const ret = await cosmJS.sendTokens(recipient, coins(1000000, "uosmo"), memo);
    console.log(ret)
}

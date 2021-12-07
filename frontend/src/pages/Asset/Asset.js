import { useState, useEffect } from 'react'
import { getBalance } from '../../helpers/getBalance';
import { Image, Table, Typography } from 'antd';
import './Asset.css'

const { Title, } = Typography;

const style = {
    divTitle: {
        display: 'flex',
        justifyContent: 'left',
        alignContent: 'left',
        textAlign: 'left',
    },
    div: {
        marginBottom: '8rem',
        borderRadius: '10px'
    }
}

const Asset = ({ address }) => {
    const [listAsset, setListAsset] = useState([]);

    useEffect(() => {
        if (address !== '') {
            (async () => {
                const balances = await getBalance(undefined, address)
                if (balances.length > 0) {
                    setListAsset([...balances])
                }
            })()
        }
        else {
            (async () => {
                const balances = await getBalance(undefined, 'osmo1cy2fkq04yh7zm6v52dm525pvx0fph7ed75lnz7')
                if (balances.length > 0) {
                    setListAsset([...balances])
                }
            })()
        }
    }, [])

    const columns = [
        {
            dataIndex: 'logo',
            key: 'logo',
            fixed: 'left',
            render: (logo) => {
                <Image
                    width={50}
                    src={logo}
                />
            }
        },
        {
            title: 'Asset/Chain',
            dataIndex: 'name',
            key: 'name',
            fixed: 'left',

        },
        {
            title: 'Amount',
            dataIndex: 'amount',
            key: 'amount',
            fixed: 'left',
        },
    ];
    return (
        <div style={style.div}>
            <div style={style.divTitle}>
                <Title>Asset List</Title>
            </div>
            <hr />
            <Table
                dataSource={listAsset}
                columns={columns}
                style={{marginTop: '3rem',  borderRadius: '5px'}}
            />
        </div>
    )
}

export default Asset;
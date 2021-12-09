
import './App.css';

import Asset from './pages/Asset/Asset'
import RootScreen from './pages/RunManuallyScreen/RootScreen';
import DepositButton from './components/DepositButton';
import DepositModal from './components/DepositModal';

import "antd/dist/antd.css";
import { Layout, Menu, Image, message, } from 'antd';
import {
  HomeOutlined,
  WalletOutlined,
  ReconciliationOutlined,
} from '@ant-design/icons';
import { Modal } from 'react-bootstrap';
import { useState, useCallback, useEffect } from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
} from "react-router-dom";

import { getKeplr, } from './helpers/getKeplr';
import { getOsmo } from './helpers/getBalance';
import logo from './assets/img/logo.png';

import { checkAccount } from './helpers/API/api';

const { Content, Sider } = Layout;

const style = {
  div: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignContent: 'center'
  },
  button: {
    borderWidth: '0px',
    borderRadius: '5px',
    size: '10em',
    backgroundColor: '#8abf80',
    color: '#ffffff',
    width: "80%",
    height: "3rem",
  },
}

function App() {
  const [account, setAccount] = useState({
    address: '',
    amount: ''
  })
  const [cookieMonster, setCookieMonster] = useState(localStorage.getItem('COOKIEMONSTER') || '')
  const [collapsed, setCollapsed] = useState(false)
  const [imageShrink, setImageShrink] = useState(false)
  const [displayTransactionModal, setDisplayTransactionModal] = useState(false)

  useEffect(() => {
    
  }, [cookieMonster])

  const wrapSetter = useCallback(value => {
    setDisplayTransactionModal(value)
  }, [setDisplayTransactionModal])

  const wrapSetCookieMonster = useCallback((value) => {
    setCookieMonster(value)
  }, [setCookieMonster])

  const wrapSetAccount = useCallback((value) => {
    setAccount({...account, amount: value})
  })

  const onCollapse = () => {
    setCollapsed(!collapsed)
    setTimeout(() => {
      setImageShrink(!imageShrink)
    }, 100);
  }

  const success = () => {
    message.success('Connect', 1);
  };

  const error = () => {
    message.error('Connect failed', 1);
  };

  const warning = () => {
    message.warning('Insufficient fund, please deposit to connect to BeanStalk', 5);
    setTimeout(() => {
      setDisplayTransactionModal(true)
    }, 1000)
  };

  const getAccount = async () => {
    const { accounts } = await getKeplr()
    const amount = await getOsmo(accounts[0].address)
    console.log(amount)
    setAccount({
      address: accounts[0].address,
      amount: amount.toString()
    })
    checkAccount(accounts[0].address).then(res => {
      if (res.data.Address !== '') {
        success()
        console.log(cookieMonster)
        setCookieMonster(res.data.Address)
        localStorage.setItem('COOKIEMONSTER', res.data.Address)
      }
      else {
        warning()
      }
    }).catch(() => {
      error()
    })
  }

  return (
    <div className="App container-fluid">
      <Layout style={{ minHeight: '100vh', marginLeft: '-12px', marginRight: '-12px', }}>
        <Router>
          <Sider theme='light'
            collapsible
            collapsed={collapsed}
            onCollapse={onCollapse}
            width={256}
            style={{ backgroundColor: '#ffffff' }}>
            <div className="logo" style={{ marginRight: '0.1rem', marginTop: '1rem', marginBottom: '1rem' }} >
              <Image
                width={!imageShrink ? 100 : 50}
                src={logo}
                preview={false}
              />
            </div>
            <hr />
            <Menu theme="light" style={{ backgroundColor: '#ffffff' }}
              mode="inline"
            >
              <Menu.Item key="home"
                icon={<HomeOutlined style={{ marginLeft: !collapsed ? '1.5rem' : '-0.3rem', fontSize: '1rem', }} />}
                style={{ margin: 0, marginTop: '10px', fontSize: '1.3rem', color: '#2b2b2b', fontWeight: 300 }}
                className="modified-item"
              >
                Home
                <Link to='/' />

              </Menu.Item>
              <Menu.Item key="asset"
                icon={<WalletOutlined style={{ marginLeft: !collapsed ? '1.5rem' : '-0.3rem', fontSize: '1rem', }} />}
                style={{ margin: 0, marginTop: '10px', fontSize: '1.3rem', color: '#2b2b2b', fontWeight: 300 }}
                className="modified-item"
              >
                Asset
                <Link to='/asset' />

              </Menu.Item>
            </Menu>
            {
              cookieMonster === '' ? (
                <div style={{ marginTop: '34rem', marginBottom: '0.3rem' }}>
                  <hr />
                  <button style={{ ...style.button, fontSize: !collapsed ? '15px' : '10px' }}
                    onClick={async () => {
                      await getAccount()
                    }}>
                    {!collapsed ? 'Connect To BeanStalk' : (<ReconciliationOutlined style={{ fontSize: '1.5rem' }} />)}
                  </button>
                </div>
              ) : (
                <DepositButton collapsed={collapsed} wrapSetter={wrapSetter} wrapSetAccount={wrapSetAccount} wrapSetCookieMonster={wrapSetCookieMonster}/>
              )
            }

          </Sider>
          <Layout className="site-layout" style={{ backgroundColor: '#c5e6be', }}>
            <Content style={{ margin: '2rem' }}>
              <div className="site-layout-background" style={{ padding: 24, paddingTop: '2rem', paddingBottom: '17rem', minHeight: 360, marginTop: '10px' }}>
                <Routes>
                  <Route exact path="/" element={<RootScreen cookieMonster={cookieMonster} account={account} />} />
                  <Route exact path="/asset" element={<Asset cookieMonster={cookieMonster} />} />
                </Routes>
              </div>
            </Content>
          </Layout>
        </Router>
      </Layout>

      <>
        <Modal show={displayTransactionModal} onHide={() => { setDisplayTransactionModal(false) }}>
          <Modal.Header closeButton>
            <Modal.Title>Deposit</Modal.Title>
          </Modal.Header>
          <Modal.Body >
            <DepositModal cookieMonster={cookieMonster} account={account} wrapSetter={wrapSetter} wrapSetCookieMonster={wrapSetCookieMonster}/>
          </Modal.Body>
        </Modal>
      </>

    </div>
  );
}

export default App;

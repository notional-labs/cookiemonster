
import './App.css';

import Register from './pages/Register';
import Asset from './pages/Asset/Asset'
import HomePage from './pages/HomePage';
import Account from './components/Account';
import RootScreen from './pages/RunManuallyScreen/RootScreen';

import "antd/dist/antd.css";
import { Layout, Menu, Image, Button } from 'antd';
import {
  HomeOutlined,
  WalletOutlined,
  UserOutlined,
  ReconciliationOutlined,
} from '@ant-design/icons';
import { useState } from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link
} from "react-router-dom";

import { getKeplr, } from './helpers/getKeplr';
import { getOsmo } from './helpers/getBalance';
import logo from './assets/img/logo.png';


const { Header, Content, Footer, Sider } = Layout;

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
    height: "2.5rem",
  },
}

function App() {
  const [account, setAccount] = useState({
    address: '',
    amount: ''
  })
  const [cookieMonster, setCookieMonster] = useState('')
  const [collapsed, setCollapsed] = useState(false)
  const [imageShrink, setImageShrink] = useState(false)

  const onCollapse = () => {
    setCollapsed(!collapsed)
    setTimeout(() => {
      setImageShrink(!imageShrink)
    }, 100);
  }

  const getAccount = async () => {
    const { accounts, offlineSigner } = await getKeplr()
    const amount = getOsmo(accounts[0].address)
    setAccount({
      address: accounts[0].address,
      amount: amount
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
                icon={<HomeOutlined style={{ marginLeft: !collapsed ? '1.5rem' : '-0.5rem', fontSize: '1.5rem', }} />}
                style={{ margin: 0, marginTop: '10px', fontSize: '1.3rem', color: '#2b2b2b', fontWeight: 300 }}
                className="modified-item"
              >
                Home
                <Link to='/' />

              </Menu.Item>
              <Menu.Item key="asset"
                icon={<WalletOutlined style={{ marginLeft: !collapsed ? '1.5rem' : '-0.5rem', fontSize: '1.5rem', }} />}
                style={{ margin: 0, marginTop: '10px', fontSize: '1.3rem', color: '#2b2b2b', fontWeight: 300 }}
                className="modified-item"
              >
                Asset
                <Link to='/asset' />

              </Menu.Item>
              <Menu.Item key="info"
                icon={<UserOutlined style={{ marginLeft: !collapsed ? '1.5rem' : '-0.5rem', fontSize: '1.5rem', }} />}
                style={{ margin: 0, marginTop: '10px', fontSize: '1.3rem', color: '#2b2b2b', fontWeight: 300 }}
                className="modified-item"
              >
                Account
                <Link to='/account' />

              </Menu.Item>
            </Menu>
            {
              cookieMonster === '' && (
                <div style={{ marginTop: '34.5rem', marginBottom: '0.3rem' }}>
                  <hr />
                  <button style={{ ...style.button, fontSize: !collapsed ? '20px' : '10px' }}
                    onClick={async () => {
                      await getAccount()
                    }}>
                    {!collapsed ? 'Connect BeanStalk' : (<ReconciliationOutlined style={{fontSize: '1.5rem'}}/>)}
                  </button>
                </div>
              ) 
            }

          </Sider>
          <Layout className="site-layout" style={{ backgroundColor: '#c5e6be', }}>
            <Content style={{ margin: '2rem' }}>
              <div className="site-layout-background" style={{ padding: 24, paddingTop: '2rem', paddingBottom: '17rem', minHeight: 360, marginTop: '10px' }}>
                <Routes>
                  <Route exact path="/" element={<RootScreen />} />
                  <Route exact path="/asset" element={<Asset address={cookieMonster} />} />
                  <Route exact path="/account" element={<Register />} />
                </Routes>
              </div>
            </Content>
          </Layout>
        </Router>
      </Layout>
    </div>
  );
}

export default App;

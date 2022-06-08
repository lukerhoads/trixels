/**
 * @type import('hardhat/config').HardhatUserConfig
 */
import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-waffle'
import '@typechain/hardhat'
import 'hardhat-abi-exporter'
import { HardhatUserConfig } from 'hardhat/types';

const config: HardhatUserConfig = {
  solidity: "0.8.6",
  typechain: {
    outDir: './client/typechain-types',
    target: 'ethers-v5',
    alwaysGenerateOverloads: false
  },
  abiExporter: [{
    path: './client/src/abi',
  }]
};

export default config
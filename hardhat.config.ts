/**
 * @type import('hardhat/config').HardhatUserConfig
 */
import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-waffle'
import '@typechain/hardhat'
import { HardhatUserConfig } from 'hardhat/types';

const config: HardhatUserConfig = {
  solidity: "0.8.6",
  typechain: {
    outDir: 'test/bindings',
    target: 'ethers-v5',
  }
};

export default config
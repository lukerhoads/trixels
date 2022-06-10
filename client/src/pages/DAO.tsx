import React, { useCallback, useEffect, useMemo, useState } from 'react';
import { useEthers } from '@usedapp/core';
import Header from 'layout/header';
import { Link } from 'react-router-dom';
import { Contract, utils } from 'ethers'
import config from '../config'

const { addresses } = config

import DAOABI from '../abi/contracts/DAO.sol/DAO.json';

const daoInterface = new utils.Interface(DAOABI);


const Dao = () => {
    const [proposals, setProposals] = useState([])
    const { library } = useEthers()

    const daoContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.dao, daoInterface, library);
    }, [library]);

    const listProposals = async () => {

    }

    const vote = useCallback(async () => {
        if (!daoContract || !library) return;
        const daoConnected = daoContract.connect(library.getSigner());
    }, [daoContract])

    return (
        <div className="dao">
            <Header>
                <Link to='/'>Home</Link>
            </Header>
            <div className="dao-wrapper">
                Dao
            </div>
        </div>
    )
}

export default Dao
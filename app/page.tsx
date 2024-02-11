"use client";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import airdrop from "@/app/airdrop"
import { useState, useEffect } from "react";
import { Connection, PublicKey, clusterApiUrl, LAMPORTS_PER_SOL, Transaction, SystemProgram, Keypair, sendAndConfirmTransaction } from '@solana/web3.js';


export default function Home() {
  const faucetAddress = process.env.NEXT_PUBLIC_FAUCET_ADDRESS;
  const airdropAmount = process.env.NEXT_PUBLIC_AIRDROP_AMOUNT;
  const [airdropResult, setAirdropResult] = useState('');
  const [faucetBalance, setFaucetBalance] = useState('');
  const [faucetEmpty, setFaucetEmpty] = useState(false);

  const handleSubmit = async (formdata: FormData) => {
    const result = await airdrop(formdata);
    setAirdropResult(result);
  }

  const getFaucetBalance = async () => {
    if(!faucetAddress) return 'No faucet!';
    const connection = new Connection(clusterApiUrl('testnet'), 'confirmed');
    const faucetPublicKey = new PublicKey(faucetAddress);
    const balanceInLamports = await connection.getBalance(faucetPublicKey);
    const balanceInSol = balanceInLamports / LAMPORTS_PER_SOL;
    if(parseInt(balanceInSol.toFixed(2)) < 2) setFaucetEmpty(true);
    return balanceInSol.toFixed(2) + ' SOL';
  }

  useEffect(() => {
    getFaucetBalance().then(balance => setFaucetBalance(balance));
  }, [airdropResult]);

  return (
    <main className="relative min-h-screen flex flex-col items-center justify-between p-4 lg:p-24">
      <header className="self-stretch flex justify-between items-center font-bold text-2xl mb-4">
        <p className="font-mono text-sm lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold">Solana</code> Testnet Faucet
        </p>
        <p className="font-mono text-sm lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold">
            <a href="https://github.com/ferric-sol/testnetfaucet">Fork on Github</a>
          </code>
        </p>
      </header>

      <form action={handleSubmit} className="flex flex-col items-center justify-center space-y-4 w-full max-w-2xl px-4">
        <div className="text-center mb-2">
          Enter wallet address to get {airdropAmount} testnet SOL airdropped
        </div>
        <div className="flex w-full">
          <input
            id="walletAddress"
            name="walletAddress"
            placeholder="Enter testnet wallet address"
            className="flex-grow px-4 py-2 border border-gray-300 rounded-l-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            onFocus={(e) => setAirdropResult('')}
            required
          />
          <button
            type="submit"
            className="px-4 py-2 bg-blue-500 text-white rounded-r-md hover:bg-blue-600 focus:ring-4 focus:ring-blue-300"
            onClick={(e) => setAirdropResult('Processing...')}
          >
            Airdrop!
          </button>
        </div>
        <p className="text-sm my-2">
          Send donation <strong>testnet</strong> sol to: {faucetAddress}
        </p>
        <p className="text-sm my-2">
          Current faucet balance is: {faucetBalance} SOL
        </p>
        <p className="text-sm my-2">
          Airdrop status: {airdropResult}
        </p>
      </form>
      <footer className="self-stretch text-center font-mono text-sm mt-4">
        Other Testnet Faucets: &nbsp;        
        <a href="https://solfaucet.com" target="_blank" rel="noopener noreferrer">SOLFaucet</a>&nbsp;
        <a href="https://faucet.quicknode.com/solana/testnet" target="_blank" rel="noopener noreferrer">Quicknode</a>&nbsp;
        <a href="https://solana.com/developers/guides/getstarted/solana-token-airdrop-and-faucets" target="_blank" rel="noopener noreferrer">Solana Guides Faucet List</a>&nbsp;
        <p className="text-xs mt-2">
          Created by <a href="https://x.com/ferric" target="_blank" rel="noopener noreferrer">@ferric</a>
        </p>
      </footer>
    </main>

  );
}
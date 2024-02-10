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
    <main className="relative min-h-screen flex flex-col items-center justify-between p-24">
      <div className="self-end font-mono text-sm">
        <p className="border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold">Solana</code>&nbsp; Testnet Faucet
        </p>
        <p className="border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold"><a href="https://github.com/ferric-sol/testnetfaucet">Fork on Github</a></code>
        </p>
      </div>

      <form action={handleSubmit} className="flex flex-col items-center justify-center space-y-4">
        <div>
          Enter wallet address to get {airdropAmount} testnet SOL airdropped
        </div>
        <div className="flex items-center space-x-2">
          <Input
            id="walletAddress"
            name="walletAddress"
            placeholder="Enter testnet wallet address"
            className="w-full px-4 py-2 border border-gray-300 rounded-l-md"
            onFocus={(e) => setAirdropResult('')}
            required
          />
          <Button
            type="submit"
            className="px-4 py-2 bg-blue-500 text-white rounded-r-md"
            onClick={(e) => setAirdropResult('Processing...')}
          >
            Airdrop!
          </Button>
        </div>
        <div className="flex items-center space-x-2">
          <p>Send donation <b>testnet</b> sol to: {faucetAddress}</p>
        </div>
        {faucetEmpty && (
          <div className="flex items-center space-x-2">
            <p>The faucet is empty! To fill up the faucet, send more <b>testnet</b> sol to: {faucetAddress}</p>
          </div>
        )}
        {!faucetEmpty && (
          <div className="flex items-center space-x-2">
            <p>Current faucet balance is: {faucetBalance}</p>
          </div>
        )}
        <div className="flex items-center space-x-2">
          <p>Airdrop status: {airdropResult}</p>
        </div>
      </form>
      <div className="flex items-center space-x-2">
        <p>
          Other Testnet Faucets:
        </p>
        <p><a href="https://solfaucet.com">SOLFaucet</a></p>
        <p><a href="https://faucet.quicknode.com/solana/testnet">Quicknode</a></p>
      </div>

      <div className="self-center w-full font-mono text-sm text-center">
        <p className="bg-transparent">
          Created by <a href="https://x.com/ferric">@ferric</a>
        </p>
      </div>
    </main>
  );
}
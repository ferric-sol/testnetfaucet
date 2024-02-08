"use client";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import airdrop from "@/app/airdrop"


export default function Home() {
  const faucetAddress = "9CfWVxa3nZwXrq2PK1czpMmJzFHmz89XpXW2cfCS3iDK";
  const faucetBalance = "100 SOL";

  return (
    <main className="relative min-h-screen flex flex-col items-center justify-between p-24">
      <div className="self-end font-mono text-sm">
        <p className="border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold">Solana</code>&nbsp; Testnet Faucet
        </p>
      </div>

      <form action={airdrop} className="flex flex-col items-center justify-center space-y-4">
        <div>
          Enter wallet address to get 1 testnet SOL airdropped
        </div>
        <div className="flex items-center space-x-2">
          <Input
            id="walletAddress"
            name="walletAddress"
            placeholder="Enter testnet wallet address"
            className="w-full px-4 py-2 border border-gray-300 rounded-l-md"
            required
          />
          <Button
            type="submit"
            className="px-4 py-2 bg-blue-500 text-white rounded-r-md"
          >
            Airdrop!
          </Button>
        </div>
        <div className="flex items-center space-x-2">
          <p>To fill up the faucet, send more <b>testnet</b> sol to: {faucetAddress}</p>
        </div>
        <div className="flex items-center space-x-2">
          <p>Current faucet balance is: {faucetBalance}</p>
        </div>

      </form>
      <div className="self-center w-full font-mono text-sm text-center">
        <p className="bg-transparent">
          Created by <a href="https://x.com/ferric">@ferric</a>
        </p>
      </div>
    </main>
  );
}
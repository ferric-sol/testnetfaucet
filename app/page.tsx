import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function Home() {
  // Placeholder content, replace with actual data fetching logic
  const faucetAddress = "9CfWVxa3nZwXrq2PK1czpMmJzFHmz89XpXW2cfCS3iDK";
  const faucetBalance = "100 SOL";

  return (
    <main className="relative min-h-screen flex flex-col items-center justify-between p-24">
      <div className="self-end font-mono text-sm">
        <p className="border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          <code className="font-bold">Solana</code>&nbsp; Testnet Faucet
        </p>
      </div>

      <div className="flex flex-col items-center justify-center space-y-4">
        <div className="relative flex place-items-center before:absolute before:h-[301px] before:w-full sm:before:w-[480px] before:-translate-x-1/2 before:rounded-full before:bg-gradient-radial before:from-white before:to-transparent before:blur-2xl before:content-[''] after:absolute after:-z-20 after:h-[180px] after:w-full sm:after:w-[240px] after:translate-x-1/3 after:bg-gradient-conic after:from-sky-200 after:via-blue-200 after:blur-2xl after:content-[''] before:dark:bg-gradient-to-br before:dark:from-transparent before:dark:to-blue-700 before:dark:opacity-10 after:dark:from-sky-900 after:dark:via-[#0141ff] after:dark:opacity-40 before:lg:h-[360px] z-[-1]">
          Enter wallet address to get 1 testnet SOL airdropped
        </div>
        <div className="relative flex place-items-center before:absolute before:h-[300px] before:w-full sm:before:w-[480px] before:-translate-x-1/2 before:rounded-full before:bg-gradient-radial before:from-white before:to-transparent before:blur-2xl before:content-[''] after:absolute after:-z-20 after:h-[180px] after:w-full sm:after:w-[240px] after:translate-x-1/3 after:bg-gradient-conic after:from-sky-200 after:via-blue-200 after:blur-2xl after:content-[''] before:dark:bg-gradient-to-br before:dark:from-transparent before:dark:to-blue-700 before:dark:opacity-10 after:dark:from-sky-900 after:dark:via-[#0141ff] after:dark:opacity-40 before:lg:h-[360px] z-[-1]">
          <Input placeholder="Enter testnet wallet address" />
          <Button type="submit">Airdrop!</Button>
        </div>
        {/* New div for displaying the faucet address and balance */}
        <div className="text-center">
          <p>To fill up the faucet, send more <b>testnet</b> sol to: {faucetAddress}</p>
          <p>Current faucet balance is: {faucetBalance}</p>
        </div>
      </div>

      <div className="self-center w-full font-mono text-sm text-center">
        <p className="bg-transparent">
          Powered by Stakeware
        </p>
      </div>
    </main>
  );
}

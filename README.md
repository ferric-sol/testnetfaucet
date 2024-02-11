# Testnet // Faucet

This app allows anyone to set up a Solana faucet. 

Fill up your NEXT_PUBLIC_FAUCET_ADDRESS with some SOL, set the secret key and the airdrop amount and you're all set!

## Getting Started

You need three environment variables in your .env.development.local file and on vercel:
- NEXT_PUBLIC_FAUCET_ADDRESS - the address of your faucet account
- SENDER_SECRET_KEY - the secret key to allow the app to send airdrops from the faucet account above
- NEXT_PUBLIC_AIRDROP_AMOUNT - the amount of SOL to send in every airdrop

You can deploy your own faucet here:

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fferric-sol%2Ftestnetfaucet&env=NEXT_PUBLIC_FAUCET_ADDRESS,SENDER_SECRET_KEY,NEXT_PUBLIC_AIRDROP_AMOUNT&envDescription=Faucet%20address%2C%20airdrop%20amount%2C%20and%20the%20faucet's%20private%20key%20are%20all%20that%20you%20need&project-name=sol-testnet-faucet&repository-name=sol-testnet-faucet&redirect-url=https%3A%2F%2Ftestnetfaucet.org&demo-title=Testnet%20Faucet&demo-description=A%20faucet%20for%20getting%20testnet%20tokens%20on%20Solana&demo-url=https%3A%2F%2Ftestnetfaucet.org&demo-image=https%3A%2F%2Fwww.stakeware.xyz%2Flogo.webp)

Or, run on localhost:

```bash
npm i 
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.


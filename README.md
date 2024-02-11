This is a Solana Faucet airdropper

## Getting Started

You need two environment variables in your .env.development.local file and on vercel:
- NEXT_PUBLIC_FAUCET_ADDRESS
- SENDER_SECRET_KEY
- NEXT_PUBLIC_AIRDROP_AMOUNT

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fferric-sol%2Ftestnetfaucet&env=NEXT_PUBLIC_FAUCET_ADDRESS,SENDER_SECRET_KEY,NEXT_PUBLIC_AIRDROP_AMOUNT&envDescription=Faucet%20address%2C%20airdrop%20amount%2C%20and%20the%20faucet's%20private%20key%20are%20all%20that%20you%20need&redirect-url=https%3A%2F%2Ftestnetfaucet.org&demo-title=Testnet%20Faucet&demo-description=A%20faucet%20for%20getting%20testnet%20tokens%20on%20Solana&demo-url=https%3A%2F%2Ftestnetfaucet.org)

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.


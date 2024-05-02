"use server";

import { Connection, PublicKey, clusterApiUrl, LAMPORTS_PER_SOL, Transaction, SystemProgram, Keypair, sendAndConfirmTransaction } from '@solana/web3.js';
import { unstable_noStore as noStore } from 'next/cache';
import { kv } from "@vercel/kv";

export default async function airdrop(formData: FormData) {
    noStore(); // Opt-into dynamic rendering

    const walletAddress = formData.get('walletAddress');
    try { 

      if (!walletAddress || walletAddress === null) {
        throw new Error('Wallet address is required');
      }

      // Connect to the cluster
      const connection = new Connection(clusterApiUrl('testnet'), 'confirmed');

      const epochInfo = await connection.getEpochInfo();

      // console.log('epochInfo: ', epochInfo);

      const walletAddressString = walletAddress?.toString();
      // Check if the wallet address is a SFDP testnet validator using
      // fetch to https://api.solana.org/api/validators/41cnC4X53ijfbYkVGoU69Qw7qRzGZ4MMbzTXvYk5xBm2
      const sfdp_url = `https://api.solana.org/api/validators/walletAddressString}`

      const sfdp_validator = await fetch(sfdp_url);
      const sfdp_validator_json = await sfdp_validator.json();

      if(!sfdp_validator_json || !sfdp_validator_json.kycStatus || !sfdp_validator_json.kycStatus as unknown === 'KYC_VALID') return 'Wallet address is not a SFDP testnet validator';

      const lastAirdropTimestampString = String(await kv.get(walletAddressString));
      const lastAirdropTimestamp = lastAirdropTimestampString ? parseInt(lastAirdropTimestampString) : null;

      const TIMEOUT_HOURS = Number(process.env.TIMEOUT_HOURS) || 24;
      const oneHourAgo = Date.now() - TIMEOUT_HOURS * 60 * 60 * 1000;

      if (lastAirdropTimestamp && lastAirdropTimestamp > oneHourAgo) {
        const minutesLeft = Math.ceil((lastAirdropTimestamp - oneHourAgo) / 60000);
        return `Try again in ${minutesLeft} minutes`;
      } 

      const secretKey = process.env.SENDER_SECRET_KEY;

      if(!secretKey) return 'Airdrop failed';

      const airdropAmount = Number(process.env.NEXT_PUBLIC_AIRDROP_AMOUNT) || 1;
      const airdropAmountLamports = airdropAmount*LAMPORTS_PER_SOL; // Send 1 SOL
      // Convert the secret key from an environment variable to a Uint8Array
      const secretKeyUint8Array = new Uint8Array(
        secretKey.split(',').map((num) => parseInt(num, 10))
      );

      // Create a keypair from the secret key
      const senderKeypair = Keypair.fromSecretKey(secretKeyUint8Array);

      // Add transfer instruction to transaction
      const transaction = new Transaction().add(
        SystemProgram.transfer({
          fromPubkey: senderKeypair.publicKey,
          toPubkey: new PublicKey(walletAddress as string),
          lamports: airdropAmountLamports
        })
      );

      // Sign the transaction with the sender's keypair
      const signature = await sendAndConfirmTransaction(
        connection,
        transaction,
        [senderKeypair]
      );

      kv.set(walletAddress as string, Date.now());

      // The transaction is now sent and confirmed, signature is the transaction id
      return 'Airdrop successful';
    } catch(error) {
      console.log('error airdropping: ', error);
      return 'Airdrop failed';
    }
}
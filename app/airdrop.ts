"use server";

export default async function airdrop(formData: FormData) {
    const walletAddress = formData.get('walletAddress');
    // Handle the response from the server action
    console.log(walletAddress);
}
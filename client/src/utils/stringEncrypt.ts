import forge from 'node-forge';

import rsa_public_keyKey from '../keys/rsa_public_key.key';

export const publicKey: string = rsa_public_keyKey;

/**
 * Encrypts data
 * @param data String to encrypt
 */
export function encrypt(data: string): string {
  const publicKeyRSA = forge.pki.publicKeyFromPem(publicKey);
  const encryptedData = publicKeyRSA.encrypt(data, 'RSAES-PKCS1-V1_5');
  return forge.util.encode64(encryptedData);
}

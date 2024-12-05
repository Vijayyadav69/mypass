// lib/encrypt_decrypt.dart
import 'dart:convert';
import 'package:encrypt/encrypt.dart';

String encrypt(String password, String masterPassword) {
  // Create a key from the master password
  final key = Key.fromUtf8(masterPassword.padRight(32, '0').substring(0, 32)); // Ensure it's 32 bytes
  final iv = IV.fromLength(16); // Random IV

  final encrypter = Encrypter(AES(key));

  // Encrypt the password
  final encrypted = encrypter.encrypt(password, iv: iv);

  // Combine the IV and the encrypted data
  return '${base64.encode(iv.bytes)}:${encrypted.base64}';
}

String decrypt(String encryptedPassword, String masterPassword) {
  // Split the IV and the encrypted data
  final parts = encryptedPassword.split(':');
  final iv = IV.fromBase64(parts[0]);
  final encryptedData = parts[1];

  // Create a key from the master password
  final key = Key.fromUtf8(masterPassword.padRight(32, '0').substring(0, 32)); // Ensure it's 32 bytes

  final encrypter = Encrypter(AES(key));

  // Decrypt the password
  return encrypter.decrypt(Encrypted.fromBase64(encryptedData), iv: iv);
}

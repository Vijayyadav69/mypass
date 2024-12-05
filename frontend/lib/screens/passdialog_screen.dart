// lib/password_dialog.dart
import 'package:flutter/material.dart';
import '../auth/encdec.dart'; // Your encryption/decryption logic

class PasswordDialog extends StatefulWidget {
  final String encryptedPassword;

  PasswordDialog(this.encryptedPassword);

  @override
  _PasswordDialogState createState() => _PasswordDialogState();
}

class _PasswordDialogState extends State<PasswordDialog> {
  final _masterPasswordController = TextEditingController();
  String? decryptedPassword;

  void _decryptPassword() {
    String masterPassword = _masterPasswordController.text;

    // Decrypt using your logic
    decryptedPassword = decrypt(widget.encryptedPassword, masterPassword);

    setState(() {}); // Refresh to show decrypted password
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text('Enter Master Password'),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          TextField(
            controller: _masterPasswordController,
            decoration: InputDecoration(labelText: 'Master Password'),
            obscureText: true,
          ),
          SizedBox(height: 20),
          ElevatedButton(
            onPressed: _decryptPassword,
            child: Text('Decrypt'),
          ),

          if (decryptedPassword != null)
            Padding(
              padding: const EdgeInsets.only(top: 20),
              child: Text('Password: $decryptedPassword    📋'),
            ),
        ],
      ),
    );
  }
}

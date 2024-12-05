import 'package:flutter/material.dart';
import '../auth/auth_service.dart';
import '../auth/user_data.dart';
import '../auth/encdec.dart'; // Assume you have encryption/decryption logic here
import 'dart:math';

class AddPasswordScreen extends StatefulWidget {
  final Function(Map<String, dynamic>) onPasswordAdded;

  AddPasswordScreen({required this.onPasswordAdded});

  @override
  _AddPasswordScreenState createState() => _AddPasswordScreenState();
}

class _AddPasswordScreenState extends State<AddPasswordScreen> {
  final _sitenameController = TextEditingController();
  final _passwordController = TextEditingController();
  final _masterPasswordController = TextEditingController();
  final AuthService _authService = AuthService();

  final _formKey = GlobalKey<FormState>(); // Global key for the form

  // Function to generate random alphanumeric password
  String _generateRandomPassword(int length) {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    Random rand = Random();
    return List.generate(length, (index) => chars[rand.nextInt(chars.length)]).join();
  }

  void _addPassword() async {
    if (_formKey.currentState?.validate() ?? false) {
      String username = UserData.username;
      String sitename = _sitenameController.text;
      String password = _passwordController.text;
      String masterPassword = _masterPasswordController.text;

      String encryptedPassword = encrypt(password, masterPassword);

      try {
        await _authService.addPassword(username, sitename, encryptedPassword);

        // Add new password to the list in HomePage
        widget.onPasswordAdded({
          'sitename': sitename,
          'password_e': encryptedPassword,
        });

        Navigator.of(context).pop(); // Go back to HomePage
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Failed to add password')));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Add Password'),
        elevation: 0,
        backgroundColor: Colors.blue,
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 24.0, vertical: 32.0),
        child: SingleChildScrollView(
          child: Form(
            key: _formKey, // Attach form key here
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Text(
                  'Add a New Password',
                  style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                  textAlign: TextAlign.center,
                ),
                SizedBox(height: 40),

                // Site Name TextField with validation
                TextFormField(
                  controller: _sitenameController,
                  decoration: InputDecoration(
                    labelText: 'Site Name',
                    border: OutlineInputBorder(),
                    prefixIcon: Icon(Icons.web),
                  ),
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return 'Please enter the site name';
                    }
                    return null;
                  },
                ),
                SizedBox(height: 16),

                // Password TextField with validation and Generate Button
                Row(
                  children: [
                    Expanded(
                      child: TextFormField(
                        controller: _passwordController,
                        decoration: InputDecoration(
                          labelText: 'Password',
                          border: OutlineInputBorder(),
                          prefixIcon: Icon(Icons.lock),
                        ),
                        obscureText: true,
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter the password';
                          }
                          return null;
                        },
                      ),
                    ),
                    IconButton(
                      icon: Icon(Icons.shuffle, color: Colors.blue),
                      onPressed: () {
                        String generatedPassword = _generateRandomPassword(12 + Random().nextInt(6)); // Random length between 10-15
                        _passwordController.text = generatedPassword;
                      },
                    ),
                  ],
                ),
                SizedBox(height: 16),

                // Master Password TextField with validation
                TextFormField(
                  controller: _masterPasswordController,
                  decoration: InputDecoration(
                    labelText: 'Master Password',
                    border: OutlineInputBorder(),
                    prefixIcon: Icon(Icons.lock_outline),
                  ),
                  obscureText: true,
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return 'Please enter the master password';
                    }
                    return null;
                  },
                ),
                SizedBox(height: 16),

                // Master Password Warning Message
                Container(
                  padding: EdgeInsets.all(12),
                  decoration: BoxDecoration(
                    color: Colors.yellow.shade100,
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Row(
                    children: [
                      Icon(
                        Icons.warning_amber_outlined,
                        color: Colors.orange,
                      ),
                      SizedBox(width: 8),
                      Expanded(
                        child: Text(
                          "Please note this master password for future use. Master password is not recoverable.",
                          style: TextStyle(
                            color: Colors.orange.shade800,
                            fontSize: 14,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                SizedBox(height: 20),

                // Add Password Button
                ElevatedButton(
                  onPressed: _addPassword,
                  style: ElevatedButton.styleFrom(
                    padding: EdgeInsets.symmetric(vertical: 16),
                    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
                    // primary: Colors.blue, // Button color
                  ),
                  child: Text(
                    'Add Password',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

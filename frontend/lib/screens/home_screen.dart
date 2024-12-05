import 'package:flutter/material.dart';
import 'package:fluttercode/screens/login_screen.dart';
import '../auth/user_data.dart';
import 'addpass_screen.dart';
import 'passdialog_screen.dart';
import '../auth/auth_service.dart';

class HomePage extends StatefulWidget {
  final List<Map<String, dynamic>>? passwordData;

  HomePage({required this.passwordData});

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final AuthService _authService = AuthService();
  void _updatePasswordList(Map<String, dynamic> newPassword) {
    setState(() {
      widget.passwordData?.add(newPassword);
    });
  }

  void _deletePassword(int index) async {
    String username = UserData.username;
    String sitename = widget.passwordData![index]['sitename'];
    String encryptedPassword = widget.passwordData![index]['password_e'];

    // Show confirmation dialog
    bool? confirmDelete = await showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: Text('Confirm Deletion'),
          content: Text('Are you sure you want to delete this password?'),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop(false); // Cancel
              },
              child: Text('Cancel'),
            ),
            TextButton(
              onPressed: () {
                Navigator.of(context).pop(true); // Confirm
              },
              child: Text('Confirm'),
            ),
          ],
        );
      },
    );

    // If the user confirmed, proceed with deletion
    if (confirmDelete ?? false) {
      try {
        // Call your delete API here
        await _authService.deletePassword(username, sitename, encryptedPassword);

        // On success, remove the password from the list
        setState(() {
          widget.passwordData?.removeAt(index);
        });

        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Password deleted successfully')));
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Failed to delete password')));
      }
    }
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        backgroundColor: Colors.transparent,
        elevation: 0,
        title: Text(
          'Passwords',
          style: TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 28,
            color: Colors.white,
          ),
        ),
        centerTitle: true,
        actions: [
          // Logout button in the top right
          IconButton(
            icon: Icon(Icons.power_settings_new_rounded,
                size: 30.0, // Increase the size for visibility
                color: Colors.black),
            onPressed: () {
              // Navigate to LoginPage or HomePage on logout
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(builder: (context) => LoginScreen()),
              );
            },
          ),
        ],
      ),
      body: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [Colors.blueAccent, Colors.lightBlueAccent],
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
          ),
        ),
        child: Column(
          children: [
            SizedBox(height: 100),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
              child: TextField(
                decoration: InputDecoration(
                  hintText: 'Search passwords...',
                  hintStyle: TextStyle(color: Colors.white70),
                  prefixIcon: Icon(Icons.search, color: Colors.white70),
                  filled: true,
                  fillColor: Colors.white24,
                  contentPadding: EdgeInsets.symmetric(vertical: 0, horizontal: 16),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(30),
                    borderSide: BorderSide.none,
                  ),
                ),
                style: TextStyle(color: Colors.white),
                onChanged: (query) {
                  // Logic to filter password list
                },
              ),
            ),
            Expanded(
              child: ListView.builder(
                padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                itemCount: widget.passwordData?.length ?? 0,
                itemBuilder: (context, index) {
                  return Padding(
                    padding: const EdgeInsets.symmetric(vertical: 6),
                    child: GestureDetector(
                      onTap: () {
                        showDialog(
                          context: context,
                          builder: (context) => PasswordDialog(widget.passwordData![index]['password_e']),
                        );
                      },
                      child: Card(
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(15),
                        ),
                        elevation: 4,
                        color: Colors.white.withOpacity(0.9),
                        child: ListTile(
                          leading: CircleAvatar(
                            backgroundColor: Colors.blueAccent,
                            child: Icon(Icons.lock_outline, color: Colors.white),
                          ),
                          title: Text(
                            widget.passwordData![index]['sitename'],
                            style: TextStyle(
                              fontWeight: FontWeight.w600,
                              fontSize: 18,
                              color: Colors.blueAccent,
                            ),
                          ),
                          trailing: Row(
                            mainAxisSize: MainAxisSize.min,
                            children: [
                              IconButton(
                                icon: Icon(Icons.delete_outline, color: Colors.red),
                                onPressed: () {
                                  // Call the delete function
                                  _deletePassword(index);
                                },
                              ),
                              Icon(Icons.arrow_forward_ios, color: Colors.grey[600], size: 18),
                            ],
                          ),
                          tileColor: Colors.white,
                        ),
                      ),
                    ),
                  );
                },
              ),
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.of(context).push(
            MaterialPageRoute(
              builder: (_) => AddPasswordScreen(onPasswordAdded: _updatePasswordList),
            ),
          );
        },
        backgroundColor: Colors.blueAccent,
        child: Icon(Icons.add, size: 30),
      ),
    );
  }
}
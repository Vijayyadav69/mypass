// lib/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;

class AuthService {
  final String baseUrl = 'http://10.0.2.2:7777/v1'; // Replace with your API base URL

  Future<List<Map<String, dynamic>>> login(String username, String password) async {
    final response = await http.post(
      Uri.parse('$baseUrl/login'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username, 'password_d': password}),
    );

    if (response.statusCode == 200) {
      // Parse the JSON response
      return List<Map<String, dynamic>>.from(jsonDecode(response.body));
    } else {
      throw Exception('Failed to login');
    }
  }

  Future<bool> register(String username, String password, String email_id) async {
    final response = await http.post(
      Uri.parse('$baseUrl/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username,
        'password_d': password,
        'email_id': email_id,
      }),
    );

    return response.statusCode == 200;
  }
  Future<void> addPassword(String username, String sitename, String password_e) async {
    final url = Uri.parse('$baseUrl/addpass');

    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
      },
      body: json.encode({
        'username': username,
        'sitename': sitename,
        'password_e': password_e,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to add password');
    }
  }

  Future<void> deletePassword(String username, String sitename, String password_e) async {
    final url = Uri.parse('$baseUrl/delpass');
    final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'username': username,
        'sitename': sitename,
        'password_e': password_e,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to delete password');
    }
  }
}

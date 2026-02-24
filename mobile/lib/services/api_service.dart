// lib/services/api_service.dart
import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:pretty_dio_logger/pretty_dio_logger.dart';
import '../models/models.dart';
import '../utils/app_theme.dart';

class ApiService {
  static final ApiService _instance = ApiService._internal();
  factory ApiService() => _instance;
  ApiService._internal();

  final _storage = const FlutterSecureStorage();
  late final Dio _dio;

  void init() {
    _dio = Dio(BaseOptions(
      baseUrl: AppConstants.baseUrl,
      connectTimeout: const Duration(seconds: 15),
      receiveTimeout: const Duration(seconds: 30),
      headers: {'Content-Type': 'application/json'},
    ));

    // Logger (debug only)
    _dio.interceptors.add(PrettyDioLogger(
      requestBody: false,
      responseBody: false,
      error: true,
    ));

    // Auth interceptor — auto-attach token & refresh on 401
    _dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) async {
        final token = await _storage.read(key: 'access_token');
        if (token != null) {
          options.headers['Authorization'] = 'Bearer $token';
        }
        handler.next(options);
      },
      onError: (error, handler) async {
        if (error.response?.statusCode == 401) {
          // Try token refresh
          final refreshed = await _refreshToken();
          if (refreshed) {
            // Retry original request
            final token = await _storage.read(key: 'access_token');
            final opts = error.requestOptions;
            opts.headers['Authorization'] = 'Bearer $token';
            try {
              final response = await _dio.fetch(opts);
              handler.resolve(response);
              return;
            } catch (_) {}
          }
          // Refresh failed → clear tokens
          await logout();
        }
        handler.next(error);
      },
    ));
  }

  Future<bool> _refreshToken() async {
    try {
      final refresh = await _storage.read(key: 'refresh_token');
      if (refresh == null) return false;
      final res = await Dio().post(
        '${AppConstants.baseUrl}/auth/refresh',
        data: {'refresh_token': refresh},
      );
      await _storage.write(key: 'access_token', value: res.data['data']['access_token']);
      await _storage.write(key: 'refresh_token', value: res.data['data']['refresh_token']);
      return true;
    } catch (_) {
      return false;
    }
  }

  // ── AUTH ──────────────────────────────────────────────────────────────────

  Future<UserModel> login(String email, String password) async {
    final res = await _dio.post('/auth/login', data: {
      'email': email,
      'password': password,
    });
    final data = res.data['data'];
    await _storage.write(key: 'access_token', value: data['access_token']);
    await _storage.write(key: 'refresh_token', value: data['refresh_token']);
    return UserModel.fromJson(data['user']);
  }

  Future<void> logout() async {
    try {
      await _dio.post('/auth/logout');
    } catch (_) {}
    await _storage.deleteAll();
  }

  Future<bool> isLoggedIn() async {
    final token = await _storage.read(key: 'access_token');
    return token != null;
  }

  Future<UserModel> getMe() async {
    final res = await _dio.get('/users/me');
    return UserModel.fromJson(res.data['data']);
  }

  // ── PORTFOLIOS ────────────────────────────────────────────────────────────

  Future<List<PortfolioModel>> getPortfolios() async {
    final res = await _dio.get('/portfolios');
    return (res.data['data'] as List)
        .map((e) => PortfolioModel.fromJson(e))
        .toList();
  }

  Future<PortfolioModel> togglePublish(String id, bool publish) async {
    final res = await _dio.patch('/portfolios/$id', data: {'is_published': publish});
    return PortfolioModel.fromJson(res.data['data']);
  }

  // ── ANALYTICS ────────────────────────────────────────────────────────────

  Future<AnalyticsSummary> getAnalytics(String portfolioId) async {
    final res = await _dio.get('/analytics/$portfolioId/summary');
    return AnalyticsSummary.fromJson(res.data['data']);
  }

  Future<AnalyticsSummary> getDashboardSummary() async {
    final res = await _dio.get('/analytics/dashboard');
    return AnalyticsSummary.fromJson(res.data['data']);
  }

  // ── NOTIFICATIONS ─────────────────────────────────────────────────────────

  Future<List<NotificationModel>> getNotifications() async {
    final res = await _dio.get('/users/notifications');
    return (res.data['data'] as List)
        .map((e) => NotificationModel.fromJson(e))
        .toList();
  }

  Future<void> markNotificationsRead() async {
    await _dio.post('/users/notifications/read-all');
  }
}

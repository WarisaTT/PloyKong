// lib/services/providers.dart
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/models.dart';
import 'api_service.dart';

final apiProvider = Provider<ApiService>((ref) {
  final api = ApiService();
  api.init();
  return api;
});

// ── Auth State ────────────────────────────────────────────────────────────────

class AuthNotifier extends AsyncNotifier<UserModel?> {
  @override
  Future<UserModel?> build() async {
    final api = ref.read(apiProvider);
    final loggedIn = await api.isLoggedIn();
    if (!loggedIn) return null;
    try {
      return await api.getMe();
    } catch (_) {
      return null;
    }
  }

  Future<void> login(String email, String password) async {
    state = const AsyncLoading();
    state = await AsyncValue.guard(() => ref.read(apiProvider).login(email, password));
  }

  Future<void> logout() async {
    await ref.read(apiProvider).logout();
    state = const AsyncData(null);
  }
}

final authProvider = AsyncNotifierProvider<AuthNotifier, UserModel?>(AuthNotifier.new);

// ── Portfolios ────────────────────────────────────────────────────────────────

final portfoliosProvider = FutureProvider<List<PortfolioModel>>((ref) async {
  return ref.read(apiProvider).getPortfolios();
});

// ── Analytics ─────────────────────────────────────────────────────────────────

final dashboardProvider = FutureProvider<AnalyticsSummary>((ref) async {
  return ref.read(apiProvider).getDashboardSummary();
});

final portfolioAnalyticsProvider =
    FutureProvider.family<AnalyticsSummary, String>((ref, id) async {
  return ref.read(apiProvider).getAnalytics(id);
});

// ── Notifications ─────────────────────────────────────────────────────────────

class NotificationsNotifier extends AsyncNotifier<List<NotificationModel>> {
  @override
  Future<List<NotificationModel>> build() async {
    return ref.read(apiProvider).getNotifications();
  }

  Future<void> markAllRead() async {
    await ref.read(apiProvider).markNotificationsRead();
    state = AsyncData(
      (state.value ?? []).map((n) => n..isRead = true).toList(),
    );
  }
}

final notificationsProvider =
    AsyncNotifierProvider<NotificationsNotifier, List<NotificationModel>>(
        NotificationsNotifier.new);

final unreadCountProvider = Provider<int>((ref) {
  return ref.watch(notificationsProvider).whenData(
        (list) => list.where((n) => !n.isRead).length,
      ).value ?? 0;
});

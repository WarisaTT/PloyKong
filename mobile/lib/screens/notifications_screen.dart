// lib/screens/notifications_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_animate/flutter_animate.dart';
import '../services/providers.dart';
import '../models/models.dart';
import '../utils/app_theme.dart';
import '../widgets/widgets.dart';
import 'package:timeago/timeago.dart' as timeago;

class NotificationsScreen extends ConsumerWidget {
  const NotificationsScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final notifAsync = ref.watch(notificationsProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: const Text('การแจ้งเตือน'),
        actions: [
          TextButton(
            onPressed: () => ref.read(notificationsProvider.notifier).markAllRead(),
            child: const Text('อ่านทั้งหมด',
                style: TextStyle(color: AppColors.neonCyan, fontSize: 13)),
          ),
        ],
      ),
      body: notifAsync.when(
        loading: () => ListView.builder(
          padding: const EdgeInsets.all(20),
          itemCount: 5,
          itemBuilder: (_, __) => Padding(
            padding: const EdgeInsets.only(bottom: 12),
            child: SkeletonBox(height: 72, radius: 16),
          ),
        ),
        error: (e, _) => Center(
          child: EmptyState(emoji: '❌', title: 'โหลดไม่ได้', subtitle: '$e'),
        ),
        data: (notifications) {
          if (notifications.isEmpty) {
            return const EmptyState(
              emoji: '🔔',
              title: 'ยังไม่มีการแจ้งเตือน',
              subtitle: 'เมื่อมีคนเข้าชมพอร์ตหรือดาวน์โหลด CV คุณจะได้รับแจ้งเตือนที่นี่',
            );
          }
          return ListView.builder(
            padding: const EdgeInsets.fromLTRB(20, 12, 20, 24),
            itemCount: notifications.length,
            itemBuilder: (ctx, i) => _NotificationTile(
              notification: notifications[i],
              delay: i * 60,
            ),
          );
        },
      ),
    );
  }
}

class _NotificationTile extends StatelessWidget {
  final NotificationModel notification;
  final int delay;

  const _NotificationTile({required this.notification, required this.delay});

  @override
  Widget build(BuildContext context) {
    final n = notification;
    final isUnread = !n.isRead;

    return Padding(
      padding: const EdgeInsets.only(bottom: 10),
      child: GlassCard(
        padding: const EdgeInsets.all(14),
        borderColor: isUnread
            ? AppColors.indigo.withOpacity(0.4)
            : AppColors.border,
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Emoji bubble
            Container(
              width: 44,
              height: 44,
              decoration: BoxDecoration(
                color: AppColors.indigo.withOpacity(0.15),
                borderRadius: BorderRadius.circular(12),
              ),
              child: Center(
                child: Text(n.emoji, style: const TextStyle(fontSize: 20)),
              ),
            ),
            const SizedBox(width: 12),

            // Content
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Expanded(
                        child: Text(
                          n.title,
                          style: TextStyle(
                            fontSize: 13,
                            fontWeight: isUnread ? FontWeight.w700 : FontWeight.w500,
                            color: isUnread ? AppColors.textPrimary : AppColors.textMuted,
                          ),
                        ),
                      ),
                      if (isUnread)
                        Container(
                          width: 8,
                          height: 8,
                          decoration: const BoxDecoration(
                            color: AppColors.neonCyan,
                            shape: BoxShape.circle,
                          ),
                        ),
                    ],
                  ),
                  const SizedBox(height: 4),
                  Text(
                    n.portfolioTitle + (n.country != null ? ' · ${n.country}' : ''),
                    style: const TextStyle(fontSize: 11, color: AppColors.textMuted),
                  ),
                  const SizedBox(height: 6),
                  Text(
                    timeago.format(n.createdAt, locale: 'th'),
                    style: const TextStyle(fontSize: 10, color: AppColors.textMuted),
                  ),
                ],
              ),
            ),
          ],
        ),
      ).animate(delay: Duration(milliseconds: delay))
          .fadeIn()
          .slideX(begin: 0.05, end: 0),
    );
  }
}

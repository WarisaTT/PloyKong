// lib/screens/portfolios_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:share_plus/share_plus.dart';
import '../services/providers.dart';
import '../models/models.dart';
import '../utils/app_theme.dart';
import '../widgets/widgets.dart';
import 'portfolio_analytics_screen.dart';
import 'package:timeago/timeago.dart' as timeago;

class PortfoliosScreen extends ConsumerWidget {
  const PortfoliosScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final portfoliosAsync = ref.watch(portfoliosProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: const Text('พอร์ตโฟลิโอ'),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: () => ref.refresh(portfoliosProvider),
          ),
        ],
      ),
      body: portfoliosAsync.when(
        loading: () => ListView.builder(
          padding: const EdgeInsets.all(20),
          itemCount: 3,
          itemBuilder: (_, i) => Padding(
            padding: const EdgeInsets.only(bottom: 16),
            child: SkeletonBox(height: 120, radius: 20),
          ),
        ),
        error: (e, _) => Center(
          child: EmptyState(
            emoji: '❌',
            title: 'โหลดข้อมูลไม่ได้',
            subtitle: '$e',
          ),
        ),
        data: (portfolios) {
          if (portfolios.isEmpty) {
            return const EmptyState(
              emoji: '🎨',
              title: 'ยังไม่มีพอร์ตโฟลิโอ',
              subtitle: 'เปิดแอป web เพื่อสร้างพอร์ตแรกของคุณ',
            );
          }
          return ListView.builder(
            padding: const EdgeInsets.fromLTRB(20, 12, 20, 24),
            itemCount: portfolios.length,
            itemBuilder: (ctx, i) => _PortfolioCard(
              portfolio: portfolios[i],
              delay: i * 100,
            ),
          );
        },
      ),
    );
  }
}

class _PortfolioCard extends ConsumerStatefulWidget {
  final PortfolioModel portfolio;
  final int delay;

  const _PortfolioCard({required this.portfolio, required this.delay});

  @override
  ConsumerState<_PortfolioCard> createState() => _PortfolioCardState();
}

class _PortfolioCardState extends ConsumerState<_PortfolioCard> {
  bool _toggling = false;

  String get _url => 'https://${widget.portfolio.slug}.ploykong.com';

  Future<void> _togglePublish() async {
    setState(() => _toggling = true);
    try {
      await ref.read(apiProvider).togglePublish(
        widget.portfolio.id,
        !widget.portfolio.isPublished,
      );
      ref.refresh(portfoliosProvider);
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Error: $e'), backgroundColor: AppColors.danger),
      );
    } finally {
      if (mounted) setState(() => _toggling = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final p = widget.portfolio;

    return Padding(
      padding: const EdgeInsets.only(bottom: 16),
      child: GlassCard(
        borderColor: p.isPublished
            ? AppColors.success.withOpacity(0.3)
            : AppColors.border,
        onTap: () => Navigator.push(
          context,
          MaterialPageRoute(
            builder: (_) => PortfolioAnalyticsScreen(portfolio: p),
          ),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Header
            Row(
              children: [
                Container(
                  width: 44,
                  height: 44,
                  decoration: BoxDecoration(
                    gradient: AppColors.gradientPrimary,
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: const Center(
                    child: Text('🌐', style: TextStyle(fontSize: 20)),
                  ),
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        p.title,
                        style: const TextStyle(
                          fontFamily: 'Syne',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                          color: AppColors.textPrimary,
                        ),
                      ),
                      Text(
                        '/${p.slug}',
                        style: const TextStyle(
                          fontSize: 12,
                          color: AppColors.textMuted,
                        ),
                      ),
                    ],
                  ),
                ),
                NeonBadge(
                  p.isPublished ? 'LIVE' : 'DRAFT',
                  color: p.isPublished ? AppColors.success : AppColors.textMuted,
                ),
              ],
            ),

            const SizedBox(height: 16),

            // Stats row
            Row(
              children: [
                _StatPill('👁️', '${p.viewCount}'),
                const SizedBox(width: 8),
                _StatPill('⏰', timeago.format(p.updatedAt)),
              ],
            ),

            const SizedBox(height: 16),
            const Divider(color: AppColors.border, height: 1),
            const SizedBox(height: 12),

            // Actions
            Row(
              children: [
                // Open in browser
                _ActionBtn(
                  icon: Icons.open_in_browser,
                  onTap: () => launchUrl(Uri.parse(_url)),
                ),
                const SizedBox(width: 8),
                // Share
                _ActionBtn(
                  icon: Icons.share_outlined,
                  onTap: () => Share.share(_url, subject: p.title),
                ),
                const SizedBox(width: 8),
                // Copy link
                _ActionBtn(
                  icon: Icons.link,
                  onTap: () {
                    Clipboard.setData(ClipboardData(text: _url));
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(
                        content: Text('คัดลอกลิงก์แล้ว! 🔗'),
                        duration: Duration(seconds: 2),
                      ),
                    );
                  },
                ),
                const Spacer(),
                // Publish toggle
                _toggling
                    ? const SizedBox(
                        width: 20,
                        height: 20,
                        child: CircularProgressIndicator(
                          strokeWidth: 2,
                          color: AppColors.neonCyan,
                        ),
                      )
                    : Switch(
                        value: p.isPublished,
                        onChanged: (_) => _togglePublish(),
                        activeColor: AppColors.success,
                        activeTrackColor: AppColors.success.withOpacity(0.3),
                        inactiveThumbColor: AppColors.textMuted,
                        inactiveTrackColor: AppColors.border,
                      ),
              ],
            ),
          ],
        ),
      ).animate(delay: Duration(milliseconds: widget.delay))
          .fadeIn()
          .slideY(begin: 0.1, end: 0),
    );
  }
}

class _StatPill extends StatelessWidget {
  final String emoji;
  final String label;
  const _StatPill(this.emoji, this.label);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
      decoration: BoxDecoration(
        color: AppColors.border.withOpacity(0.5),
        borderRadius: BorderRadius.circular(20),
      ),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(emoji, style: const TextStyle(fontSize: 12)),
          const SizedBox(width: 4),
          Text(label, style: const TextStyle(
            fontSize: 11,
            color: AppColors.textMuted,
          )),
        ],
      ),
    );
  }
}

class _ActionBtn extends StatelessWidget {
  final IconData icon;
  final VoidCallback onTap;
  const _ActionBtn({required this.icon, required this.onTap});

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      borderRadius: BorderRadius.circular(10),
      child: Container(
        width: 36,
        height: 36,
        decoration: BoxDecoration(
          color: AppColors.border.withOpacity(0.5),
          borderRadius: BorderRadius.circular(10),
        ),
        child: Icon(icon, size: 18, color: AppColors.textMuted),
      ),
    );
  }
}

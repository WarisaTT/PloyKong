// lib/screens/profile_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:url_launcher/url_launcher.dart';
import '../services/providers.dart';
import '../utils/app_theme.dart';
import '../widgets/widgets.dart';

class ProfileScreen extends ConsumerWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authAsync = ref.watch(authProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(title: const Text('โปรไฟล์')),
      body: authAsync.when(
        loading: () => const Center(child: CircularProgressIndicator(color: AppColors.neonCyan)),
        error: (e, _) => Center(child: EmptyState(emoji: '❌', title: '$e')),
        data: (user) {
          if (user == null) return const SizedBox.shrink();
          return SingleChildScrollView(
            padding: const EdgeInsets.all(20),
            child: Column(
              children: [
                // Avatar
                Container(
                  width: 80,
                  height: 80,
                  decoration: BoxDecoration(
                    gradient: AppColors.gradientPrimary,
                    shape: BoxShape.circle,
                    boxShadow: [
                      BoxShadow(
                        color: AppColors.indigo.withOpacity(0.4),
                        blurRadius: 20,
                        spreadRadius: -5,
                      ),
                    ],
                  ),
                  child: Center(
                    child: Text(
                      user.name.substring(0, 1).toUpperCase(),
                      style: const TextStyle(
                        fontFamily: 'Syne',
                        fontSize: 32,
                        fontWeight: FontWeight.w800,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 16),
                Text(user.name, style: const TextStyle(
                  fontFamily: 'Syne',
                  fontSize: 22,
                  fontWeight: FontWeight.w800,
                  color: AppColors.textPrimary,
                )),
                const SizedBox(height: 4),
                Text(user.email, style: const TextStyle(
                  fontSize: 14,
                  color: AppColors.textMuted,
                )),
                const SizedBox(height: 8),
                NeonBadge(
                  user.plan.toUpperCase(),
                  color: user.plan == 'pro' ? AppColors.neonPurple : AppColors.textMuted,
                ),

                const SizedBox(height: 32),

                // Menu items
                GlassCard(
                  padding: EdgeInsets.zero,
                  child: Column(
                    children: [
                      _MenuItem(
                        icon: Icons.web_outlined,
                        label: 'เปิดเว็บ Builder',
                        onTap: () => launchUrl(Uri.parse('https://app.ploykong.com')),
                      ),
                      const Divider(color: AppColors.border, height: 1),
                      _MenuItem(
                        icon: Icons.star_outline,
                        label: 'อัปเกรดเป็น Pro',
                        trailing: NeonBadge('PRO', color: AppColors.neonPurple),
                        onTap: () => launchUrl(Uri.parse('https://app.ploykong.com/pricing')),
                      ),
                      const Divider(color: AppColors.border, height: 1),
                      _MenuItem(
                        icon: Icons.help_outline,
                        label: 'ช่วยเหลือ & FAQ',
                        onTap: () => launchUrl(Uri.parse('https://app.ploykong.com/help')),
                      ),
                      const Divider(color: AppColors.border, height: 1),
                      _MenuItem(
                        icon: Icons.info_outline,
                        label: 'เวอร์ชัน 1.0.0',
                        onTap: null,
                        trailing: const Text('PloyKong Mobile',
                            style: TextStyle(fontSize: 12, color: AppColors.textMuted)),
                      ),
                    ],
                  ),
                ),

                const SizedBox(height: 20),

                // Logout
                NeonButton(
                  text: 'ออกจากระบบ',
                  gradient: LinearGradient(
                    colors: [AppColors.danger.withOpacity(0.8), AppColors.danger],
                  ),
                  onPressed: () async {
                    await ref.read(authProvider.notifier).logout();
                    if (context.mounted) {
                      Navigator.pushReplacementNamed(context, '/login');
                    }
                  },
                  width: double.infinity,
                ),

                const SizedBox(height: 32),

                const Text(
                  'PloyKong © 2025 · KMUTT IT',
                  style: TextStyle(fontSize: 11, color: AppColors.textMuted),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}

class _MenuItem extends StatelessWidget {
  final IconData icon;
  final String label;
  final VoidCallback? onTap;
  final Widget? trailing;

  const _MenuItem({
    required this.icon,
    required this.label,
    this.onTap,
    this.trailing,
  });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      onTap: onTap,
      leading: Icon(icon, color: AppColors.textMuted, size: 20),
      title: Text(label, style: const TextStyle(
        fontSize: 14,
        color: AppColors.textPrimary,
        fontWeight: FontWeight.w500,
      )),
      trailing: trailing ?? (onTap != null
          ? const Icon(Icons.chevron_right, color: AppColors.textMuted, size: 18)
          : null),
    );
  }
}

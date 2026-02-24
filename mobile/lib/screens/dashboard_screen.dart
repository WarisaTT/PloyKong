// lib/screens/dashboard_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:flutter_animate/flutter_animate.dart';
import '../services/providers.dart';
import '../models/models.dart';
import '../utils/app_theme.dart';
import '../widgets/widgets.dart';
import 'package:intl/intl.dart';

class DashboardScreen extends ConsumerWidget {
  const DashboardScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final auth = ref.watch(authProvider);
    final dashAsync = ref.watch(dashboardProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      body: CustomScrollView(
        slivers: [
          // ── App Bar ─────────────────────────────────────────────────────────
          SliverAppBar(
            expandedHeight: 120,
            floating: true,
            snap: true,
            backgroundColor: AppColors.bg,
            elevation: 0,
            flexibleSpace: FlexibleSpaceBar(
              background: Padding(
                padding: const EdgeInsets.fromLTRB(20, 60, 20, 0),
                child: auth.when(
                  data: (user) => Row(
                    children: [
                      CircleAvatar(
                        radius: 24,
                        backgroundColor: AppColors.indigo,
                        child: Text(
                          user?.name.substring(0, 1).toUpperCase() ?? 'P',
                          style: const TextStyle(
                            fontFamily: 'Syne',
                            fontSize: 20,
                            fontWeight: FontWeight.w800,
                            color: Colors.white,
                          ),
                        ),
                      ),
                      const SizedBox(width: 14),
                      Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          const Text('สวัสดี 👋', style: TextStyle(
                            fontSize: 13,
                            color: AppColors.textMuted,
                          )),
                          Text(
                            user?.name ?? 'Ployer',
                            style: const TextStyle(
                              fontFamily: 'Syne',
                              fontSize: 18,
                              fontWeight: FontWeight.w700,
                              color: AppColors.textPrimary,
                            ),
                          ),
                        ],
                      ),
                      const Spacer(),
                      NeonBadge(
                        auth.value?.plan.toUpperCase() ?? 'FREE',
                        color: AppColors.neonPurple,
                      ),
                    ],
                  ),
                  loading: () => const SizedBox.shrink(),
                  error: (_, __) => const SizedBox.shrink(),
                ),
              ),
            ),
          ),

          // ── Content ─────────────────────────────────────────────────────────
          SliverPadding(
            padding: const EdgeInsets.fromLTRB(20, 8, 20, 24),
            sliver: dashAsync.when(
              loading: () => SliverList(
                delegate: SliverChildListDelegate([
                  const SizedBox(height: 16),
                  const SkeletonBox(height: 100),
                  const SizedBox(height: 16),
                  const SkeletonBox(height: 200),
                  const SizedBox(height: 16),
                  const SkeletonBox(height: 150),
                ]),
              ),
              error: (e, _) => SliverFillRemaining(
                child: Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text('❌', style: TextStyle(fontSize: 40)),
                      const SizedBox(height: 16),
                      Text('$e', style: const TextStyle(color: AppColors.textMuted)),
                      const SizedBox(height: 16),
                      ElevatedButton(
                        onPressed: () => ref.refresh(dashboardProvider),
                        child: const Text('ลองใหม่'),
                      ),
                    ],
                  ),
                ),
              ),
              data: (summary) => _DashboardContent(summary: summary),
            ),
          ),
        ],
      ),
    );
  }
}

class _DashboardContent extends StatelessWidget {
  final AnalyticsSummary summary;
  const _DashboardContent({required this.summary});

  @override
  Widget build(BuildContext context) {
    return SliverList(
      delegate: SliverChildListDelegate([
        // ── Stat Cards Grid ──────────────────────────────────────────────────
        GridView.count(
          crossAxisCount: 2,
          crossAxisSpacing: 12,
          mainAxisSpacing: 12,
          shrinkWrap: true,
          physics: const NeverScrollableScrollPhysics(),
          childAspectRatio: 1.3,
          children: [
            StatCard(
              emoji: '👁️',
              label: 'Total Views',
              value: _fmt(summary.totalViews),
              sub: '+${summary.todayViews} วันนี้',
              color: AppColors.neonCyan,
              delay: 0,
            ),
            StatCard(
              emoji: '📄',
              label: 'PDF Downloads',
              value: _fmt(summary.pdfDownloads),
              color: AppColors.neonPurple,
              delay: 100,
            ),
            StatCard(
              emoji: '🎯',
              label: 'Hire Me',
              value: _fmt(summary.hireClicks),
              color: AppColors.hotPink,
              delay: 200,
            ),
            StatCard(
              emoji: '🤖',
              label: 'AI Chats',
              value: _fmt(summary.aiChatCount),
              color: AppColors.success,
              delay: 300,
            ),
          ],
        ),

        const SizedBox(height: 24),

        // ── Bar Chart ────────────────────────────────────────────────────────
        if (summary.dailyViews.isNotEmpty) ...[
          GlassCard(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text(
                  '📈 Views รายวัน',
                  style: TextStyle(
                    fontFamily: 'Syne',
                    fontSize: 16,
                    fontWeight: FontWeight.w700,
                    color: AppColors.textPrimary,
                  ),
                ),
                const SizedBox(height: 4),
                Text(
                  '7 วันล่าสุด · ${summary.weekViews} views รวม',
                  style: const TextStyle(fontSize: 12, color: AppColors.textMuted),
                ),
                const SizedBox(height: 20),
                SizedBox(
                  height: 140,
                  child: BarChart(
                    BarChartData(
                      alignment: BarChartAlignment.spaceAround,
                      maxY: summary.dailyViews
                          .map((d) => d.count.toDouble())
                          .fold(0.0, (a, b) => a > b ? a : b) *
                          1.3,
                      barGroups: summary.dailyViews.asMap().entries.map((e) {
                        return BarChartGroupData(
                          x: e.key,
                          barRods: [
                            BarChartRodData(
                              toY: e.value.count.toDouble(),
                              gradient: const LinearGradient(
                                colors: [AppColors.indigo, AppColors.neonPurple],
                                begin: Alignment.bottomCenter,
                                end: Alignment.topCenter,
                              ),
                              width: 22,
                              borderRadius: const BorderRadius.vertical(
                                top: Radius.circular(6),
                              ),
                            ),
                          ],
                        );
                      }).toList(),
                      titlesData: FlTitlesData(
                        leftTitles: const AxisTitles(
                          sideTitles: SideTitles(showTitles: false),
                        ),
                        rightTitles: const AxisTitles(
                          sideTitles: SideTitles(showTitles: false),
                        ),
                        topTitles: const AxisTitles(
                          sideTitles: SideTitles(showTitles: false),
                        ),
                        bottomTitles: AxisTitles(
                          sideTitles: SideTitles(
                            showTitles: true,
                            getTitlesWidget: (value, meta) {
                              final idx = value.toInt();
                              if (idx >= 0 && idx < summary.dailyViews.length) {
                                final date = summary.dailyViews[idx].date;
                                final formatted = date.length >= 10
                                    ? date.substring(5) // MM-DD
                                    : date;
                                return Text(
                                  formatted,
                                  style: const TextStyle(
                                    fontSize: 9,
                                    color: AppColors.textMuted,
                                  ),
                                );
                              }
                              return const SizedBox.shrink();
                            },
                          ),
                        ),
                      ),
                      gridData: FlGridData(
                        show: true,
                        drawVerticalLine: false,
                        getDrawingHorizontalLine: (value) => FlLine(
                          color: AppColors.border,
                          strokeWidth: 1,
                          dashArray: [4, 4],
                        ),
                      ),
                      borderData: FlBorderData(show: false),
                      barTouchData: BarTouchData(
                        touchTooltipData: BarTouchTooltipData(
                          getTooltipColor: (_) => AppColors.surface,
                          tooltipRoundedRadius: 8,
                          getTooltipItem: (group, groupIndex, rod, rodIndex) {
                            return BarTooltipItem(
                              '${rod.toY.toInt()} views',
                              const TextStyle(
                                color: AppColors.neonCyan,
                                fontWeight: FontWeight.w700,
                                fontSize: 12,
                              ),
                            );
                          },
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ).animate().fadeIn(delay: 400.ms),

          const SizedBox(height: 16),
        ],

        // ── Source Breakdown (Pie) ────────────────────────────────────────────
        if (summary.sourceStats.isNotEmpty) ...[
          GlassCard(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text(
                  '🔗 แหล่งที่มาของผู้เข้าชม',
                  style: TextStyle(
                    fontFamily: 'Syne',
                    fontSize: 16,
                    fontWeight: FontWeight.w700,
                    color: AppColors.textPrimary,
                  ),
                ),
                const SizedBox(height: 20),
                Row(
                  children: [
                    SizedBox(
                      width: 120,
                      height: 120,
                      child: PieChart(
                        PieChartData(
                          sections: _buildPieSections(summary.sourceStats),
                          centerSpaceRadius: 35,
                          sectionsSpace: 2,
                        ),
                      ),
                    ),
                    const SizedBox(width: 20),
                    Expanded(
                      child: Column(
                        children: summary.sourceStats.asMap().entries.map((e) {
                          final colors = [
                            AppColors.indigo,
                            AppColors.neonCyan,
                            AppColors.neonPurple,
                            AppColors.hotPink,
                            AppColors.success,
                          ];
                          final color = colors[e.key % colors.length];
                          final total = summary.sourceStats
                              .fold(0, (s, x) => s + x.count);
                          final pct = total > 0
                              ? (e.value.count / total * 100).toStringAsFixed(0)
                              : '0';
                          return Padding(
                            padding: const EdgeInsets.only(bottom: 8),
                            child: Row(
                              children: [
                                Container(
                                  width: 10,
                                  height: 10,
                                  decoration: BoxDecoration(
                                    color: color,
                                    borderRadius: BorderRadius.circular(3),
                                  ),
                                ),
                                const SizedBox(width: 8),
                                Expanded(
                                  child: Text(
                                    e.value.source,
                                    style: const TextStyle(
                                      fontSize: 12,
                                      color: AppColors.textPrimary,
                                    ),
                                    overflow: TextOverflow.ellipsis,
                                  ),
                                ),
                                Text(
                                  '$pct%',
                                  style: TextStyle(
                                    fontSize: 12,
                                    color: color,
                                    fontWeight: FontWeight.w700,
                                  ),
                                ),
                              ],
                            ),
                          );
                        }).toList(),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ).animate().fadeIn(delay: 500.ms),

          const SizedBox(height: 16),
        ],

        // ── Country List ─────────────────────────────────────────────────────
        if (summary.countryStats.isNotEmpty) ...[
          GlassCard(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text(
                  '🌍 ผู้เข้าชมตามประเทศ',
                  style: TextStyle(
                    fontFamily: 'Syne',
                    fontSize: 16,
                    fontWeight: FontWeight.w700,
                    color: AppColors.textPrimary,
                  ),
                ),
                const SizedBox(height: 16),
                ...summary.countryStats.take(5).map((c) {
                  final pct = summary.totalViews > 0
                      ? c.count / summary.totalViews
                      : 0.0;
                  return Padding(
                    padding: const EdgeInsets.only(bottom: 12),
                    child: Row(
                      children: [
                        Text(
                          _flagEmoji(c.countryCode),
                          style: const TextStyle(fontSize: 20),
                        ),
                        const SizedBox(width: 10),
                        SizedBox(
                          width: 40,
                          child: Text(
                            c.countryCode,
                            style: const TextStyle(
                              fontSize: 12,
                              fontWeight: FontWeight.w600,
                              color: AppColors.textPrimary,
                            ),
                          ),
                        ),
                        Expanded(
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(4),
                            child: LinearProgressIndicator(
                              value: pct,
                              minHeight: 6,
                              backgroundColor: AppColors.border,
                              valueColor: const AlwaysStoppedAnimation(
                                AppColors.neonCyan,
                              ),
                            ),
                          ),
                        ),
                        const SizedBox(width: 10),
                        Text(
                          '${c.count}',
                          style: const TextStyle(
                            fontSize: 12,
                            color: AppColors.textMuted,
                          ),
                        ),
                      ],
                    ),
                  );
                }),
              ],
            ),
          ).animate().fadeIn(delay: 600.ms),
        ],
      ]),
    );
  }

  List<PieChartSectionData> _buildPieSections(List<SourceStat> stats) {
    final colors = [
      AppColors.indigo,
      AppColors.neonCyan,
      AppColors.neonPurple,
      AppColors.hotPink,
      AppColors.success,
    ];
    final total = stats.fold(0, (s, x) => s + x.count);
    return stats.asMap().entries.map((e) {
      return PieChartSectionData(
        color: colors[e.key % colors.length],
        value: e.value.count.toDouble(),
        title: total > 0
            ? '${(e.value.count / total * 100).toStringAsFixed(0)}%'
            : '',
        radius: 40,
        titleStyle: const TextStyle(
          fontSize: 10,
          fontWeight: FontWeight.w700,
          color: Colors.white,
        ),
      );
    }).toList();
  }

  String _flagEmoji(String code) {
    if (code.length != 2) return '🌐';
    final upper = code.toUpperCase();
    final first = String.fromCharCode(0x1F1E6 - 65 + upper.codeUnitAt(0));
    final second = String.fromCharCode(0x1F1E6 - 65 + upper.codeUnitAt(1));
    return '$first$second';
  }

  String _fmt(int n) => NumberFormat.compact().format(n);
}

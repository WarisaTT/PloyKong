// lib/screens/portfolio_analytics_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:fl_chart/fl_chart.dart';
import '../models/models.dart';
import '../services/providers.dart';
import '../utils/app_theme.dart';
import '../widgets/widgets.dart';

class PortfolioAnalyticsScreen extends ConsumerWidget {
  final PortfolioModel portfolio;
  const PortfolioAnalyticsScreen({super.key, required this.portfolio});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final analyticsAsync = ref.watch(portfolioAnalyticsProvider(portfolio.id));

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(portfolio.title,
                style: const TextStyle(fontSize: 16, fontWeight: FontWeight.w700)),
            Text('/${portfolio.slug}',
                style: const TextStyle(fontSize: 11, color: AppColors.textMuted)),
          ],
        ),
      ),
      body: analyticsAsync.when(
        loading: () => const Center(child: CircularProgressIndicator(
          color: AppColors.neonCyan,
        )),
        error: (e, _) => Center(child: EmptyState(
          emoji: '❌', title: 'โหลดไม่ได้', subtitle: '$e',
        )),
        data: (summary) => SingleChildScrollView(
          padding: const EdgeInsets.all(20),
          child: Column(
            children: [
              // Stat cards
              GridView.count(
                crossAxisCount: 2,
                crossAxisSpacing: 12,
                mainAxisSpacing: 12,
                shrinkWrap: true,
                physics: const NeverScrollableScrollPhysics(),
                childAspectRatio: 1.3,
                children: [
                  StatCard(
                    emoji: '👁️', label: 'Views ทั้งหมด',
                    value: '${summary.totalViews}',
                    sub: '${summary.todayViews} วันนี้',
                    color: AppColors.neonCyan,
                  ),
                  StatCard(
                    emoji: '📄', label: 'PDF',
                    value: '${summary.pdfDownloads}',
                    color: AppColors.neonPurple,
                  ),
                  StatCard(
                    emoji: '🎯', label: 'Hire Me',
                    value: '${summary.hireClicks}',
                    color: AppColors.hotPink,
                  ),
                  StatCard(
                    emoji: '🤖', label: 'AI Chat',
                    value: '${summary.aiChatCount}',
                    color: AppColors.success,
                  ),
                ],
              ),

              const SizedBox(height: 20),

              // Bar chart
              if (summary.dailyViews.isNotEmpty)
                GlassCard(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text('📈 Views 7 วัน', style: TextStyle(
                        fontFamily: 'Syne', fontSize: 16,
                        fontWeight: FontWeight.w700, color: AppColors.textPrimary,
                      )),
                      const SizedBox(height: 20),
                      SizedBox(
                        height: 150,
                        child: BarChart(BarChartData(
                          alignment: BarChartAlignment.spaceAround,
                          maxY: summary.dailyViews
                                  .map((d) => d.count.toDouble())
                                  .fold(0.0, (a, b) => a > b ? a : b) * 1.3 + 1,
                          barGroups: summary.dailyViews.asMap().entries.map((e) =>
                            BarChartGroupData(x: e.key, barRods: [
                              BarChartRodData(
                                toY: e.value.count.toDouble(),
                                gradient: const LinearGradient(
                                  colors: [AppColors.indigo, AppColors.neonPurple],
                                  begin: Alignment.bottomCenter,
                                  end: Alignment.topCenter,
                                ),
                                width: 20,
                                borderRadius: const BorderRadius.vertical(
                                  top: Radius.circular(6),
                                ),
                              ),
                            ]),
                          ).toList(),
                          titlesData: FlTitlesData(
                            leftTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                            rightTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                            topTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                            bottomTitles: AxisTitles(
                              sideTitles: SideTitles(
                                showTitles: true,
                                getTitlesWidget: (v, _) {
                                  final i = v.toInt();
                                  if (i >= 0 && i < summary.dailyViews.length) {
                                    return Text(
                                      summary.dailyViews[i].date.length >= 10
                                          ? summary.dailyViews[i].date.substring(5)
                                          : summary.dailyViews[i].date,
                                      style: const TextStyle(fontSize: 9, color: AppColors.textMuted),
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
                            getDrawingHorizontalLine: (_) => FlLine(
                              color: AppColors.border,
                              strokeWidth: 1,
                              dashArray: [4, 4],
                            ),
                          ),
                          borderData: FlBorderData(show: false),
                        )),
                      ),
                    ],
                  ),
                ),

              const SizedBox(height: 16),

              // Country list
              if (summary.countryStats.isNotEmpty)
                GlassCard(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text('🌍 ประเทศ', style: TextStyle(
                        fontFamily: 'Syne', fontSize: 16,
                        fontWeight: FontWeight.w700, color: AppColors.textPrimary,
                      )),
                      const SizedBox(height: 16),
                      ...summary.countryStats.take(8).map((c) {
                        final pct = summary.totalViews > 0
                            ? c.count / summary.totalViews
                            : 0.0;
                        return Padding(
                          padding: const EdgeInsets.only(bottom: 12),
                          child: Row(
                            children: [
                              Text(_flagEmoji(c.countryCode),
                                  style: const TextStyle(fontSize: 18)),
                              const SizedBox(width: 8),
                              SizedBox(
                                width: 36,
                                child: Text(c.countryCode, style: const TextStyle(
                                  fontSize: 12, color: AppColors.textPrimary,
                                  fontWeight: FontWeight.w600,
                                )),
                              ),
                              Expanded(
                                child: ClipRRect(
                                  borderRadius: BorderRadius.circular(4),
                                  child: LinearProgressIndicator(
                                    value: pct.toDouble(),
                                    minHeight: 6,
                                    backgroundColor: AppColors.border,
                                    valueColor: const AlwaysStoppedAnimation(AppColors.indigo),
                                  ),
                                ),
                              ),
                              const SizedBox(width: 8),
                              Text('${c.count}', style: const TextStyle(
                                fontSize: 12, color: AppColors.textMuted,
                              )),
                            ],
                          ),
                        );
                      }),
                    ],
                  ),
                ),
            ],
          ),
        ),
      ),
    );
  }

  String _flagEmoji(String code) {
    if (code.length != 2) return '🌐';
    final upper = code.toUpperCase();
    final first = String.fromCharCode(0x1F1E6 - 65 + upper.codeUnitAt(0));
    final second = String.fromCharCode(0x1F1E6 - 65 + upper.codeUnitAt(1));
    return '$first$second';
  }
}

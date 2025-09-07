'use client';

import { useTransition } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/component/base/Button';
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '@/component/base/Form';
import { Input } from '@/component/base/Input';
import { Text, Title } from '@/component/base/Typography';
import { setUserLocale } from '@/i18n';
import { Locale } from '@/i18n/conf';
import { zodResolver } from '@hookform/resolvers/zod';
import { useTranslations } from 'next-intl';
import { useTheme } from 'next-themes';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

const formSchema = z.object({
	username: z.string().min(2, {
		message: 'Username must be at least 2 characters.',
	}),
});

export default function Page() {
	const t = useTranslations('Login');
	const { setTheme } = useTheme();
	const [isPending, startTransition] = useTransition();
	const router = useRouter();

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			username: '',
		},
	});

	// 2. Define a submit handler.
	function onSubmit(values: z.infer<typeof formSchema>) {
		// Do something with the form values.
		// ✅ This will be type-safe and validated.
		console.log(values);
	}

	const handlerThemeChang = (e: string) => {
		setTheme(e);
	};

	const handleri18nChang = (e: Locale) => {
		setUserLocale(e);
	};

	return (
		<div className='flex flex-col md:w-[400px]'>
			<Title size='h2'>{t('title')}</Title>
			<Text>{t('subtitle', { name: 'Nextadm' })}</Text>
			<Form {...form}>
				<form onSubmit={form.handleSubmit(onSubmit)} className='space-y-8'>
					<FormField
						control={form.control}
						name='username'
						render={({ field }) => (
							<FormItem>
								<FormLabel>Username</FormLabel>
								<FormControl>
									<Input placeholder='shadcn' {...field} />
								</FormControl>
								<FormDescription>This is your public display name.</FormDescription>
								<FormMessage />
							</FormItem>
						)}
					/>
					<Button type='submit'>登 录</Button>
				</form>
			</Form>

			<Button onClick={(e) => handlerThemeChang('dark')}>测试 dark</Button>
			<Button onClick={(e) => handlerThemeChang('light')}>测试 light</Button>

			<Button onClick={(e) => handleri18nChang('en')}>测试 i18n es-US</Button>
			<Button onClick={(e) => handleri18nChang('zh')}>测试 i18n zh-Hant</Button>
		</div>
	);
}
